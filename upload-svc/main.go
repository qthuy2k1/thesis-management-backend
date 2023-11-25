package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	uploadpb "github.com/qthuy2k1/thesis-management-backend/upload-svc/api/goclient/v1"
)

func main() {
	// create new router
	r := chi.NewRouter()

	client := NewUploadService()

	r.Route("/api/exercise", func(r chi.Router) {
		r.Post("/", client.createExercise)
	})

	log.Println("Upload service starting on 0.0.0.0:8081")
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", 8081), r); err != nil {
		log.Fatal(err)
	}
}

type FileInfo struct {
	FileName  string
	Thumbnail string
	Size      int64
	MimeType  string
	URL       string
}

func NewUploadService() *uploadServiceGW {
	// Specify the address and port of the gRPC gateway
	gatewayAddress := "thesis-management-backend-service:9091"

	// Create a connection to the gRPC gateway
	conn, err := grpc.Dial(gatewayAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC gateway: %v", err)
	}

	client := uploadpb.NewUploadServiceClient(conn)
	exerciseClient := pb.NewExerciseServiceClient(conn)

	return &uploadServiceGW{
		client: client,
		gw:     exerciseClient,
	}
}

type uploadServiceGW struct {
	uploadpb.UnimplementedUploadServiceServer
	client uploadpb.UploadServiceClient
	gw     pb.ExerciseServiceClient
}

func (u *uploadServiceGW) createExercise(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: createExercise is called")
	// Parse the multipart form data
	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	title := r.FormValue("title")
	description := r.FormValue("description")
	classroomID := r.FormValue("classroomID")
	deadline := r.FormValue("deadline")
	categoryID := r.FormValue("categoryID")
	authorID := r.FormValue("authorID")

	// attachment
	var attachments []*pb.AttachmentExerciseInput
	fhs := r.MultipartForm.File["attachments"]
	for _, fileHeader := range fhs {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		driveFile, err := uploadFileToDrive(ctx, file, fileHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileInfo := FileInfo{
			FileName:  driveFile.Name,
			Thumbnail: driveFile.ThumbnailLink,
			Size:      driveFile.Size,
			MimeType:  driveFile.MimeType,
			URL:       driveFile.WebViewLink,
		}

		attachments = append(attachments, &pb.AttachmentExerciseInput{
			FileURL:   fileInfo.URL,
			AuthorID:  authorID,
			Name:      fileInfo.FileName,
			Size:      fileInfo.Size,
			Type:      fileInfo.MimeType,
			Thumbnail: fileInfo.Thumbnail,
		})
	}

	classroomIDInt, err := strconv.Atoi(classroomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02 15:04:05", deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.gw.CreateExercise(ctx, &pb.CreateExerciseRequest{
		Exercise: &pb.ExerciseInput{
			Title:       title,
			Description: description,
			ClassroomID: int64(classroomIDInt),
			Deadline: &datetime.DateTime{
				Year:    int32(date.Year()),
				Month:   int32(date.Month()),
				Day:     int32(date.Day()),
				Hours:   int32(date.Hour()),
				Minutes: int32(date.Minute()),
				Seconds: int32(date.Second()),
			},
			CategoryID:  int64(categoryIDInt),
			AuthorID:    authorID,
			Attachments: attachments,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"code":    res.Response.StatusCode,
		"message": res.Response.Message,
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Response.StatusCode))
	w.Write(jsonBytes)
}

func uploadFileToDrive(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*drive.File, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, err
	}
	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	// Create a new file metadata
	driveFile := &drive.File{
		Name: header.Filename,
	}

	// Upload the file
	driveFile, err = srv.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		return nil, err
	}

	if _, err := srv.Permissions.Create(driveFile.Id, &drive.Permission{
		Role: "reader",
		Type: "anyone",
	}).Do(); err != nil {
		return nil, err
	}

	getDriveFile, err := srv.Files.Get(driveFile.Id).Fields("name, size, thumbnailLink, mimeType, webViewLink").Do()
	if err != nil {
		return nil, err
	}

	return getDriveFile, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	log.Println("Input here")
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
