package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang/glog"
	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"

	// uploadpb "github.com/qthuy2k1/thesis-management-backend/upload-svc/api/goclient/v1"

	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
)

type FileInfo struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Size      int64  `json:"size"`
	Type      string `json:"type"`
	FileURL   string `json:"fileURL"`
	AuthorID  string `json:"authorID"`
	Status    string `json:"status"`
}

type Response struct {
	Response BodyResponse `json:"response"`
}

type BodyResponse struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		preflightHandler(w, r)
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}

func main() {
	// create new router
	r := chi.NewRouter()

	client := NewUploadService()

	r.Use(allowCORS)

	r.Use(middleware.Logger)

	r.Route("/upload", func(r chi.Router) {
		r.Use(allowCORS)
		r.Route("/exercise", func(r chi.Router) {
			r.Post("/", client.createExercise)
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Test get server"))
			})
			r.Route("/{exerciseID}", func(r chi.Router) {
				r.Put("/", client.updateExercise)
			})
		})
		r.Route("/post", func(r chi.Router) {
			r.Post("/", client.createPost)
			r.Route("/{postID}", func(r chi.Router) {
				r.Put("/", client.updatePost)
			})
		})
		r.Route("/submit", func(r chi.Router) {
			r.Post("/", client.createSubmission)
			r.Route("/{submissionID}", func(r chi.Router) {
				r.Put("/", client.updateSubmission)
			})
		})

		r.Route("/final-file", func(r chi.Router) {
			r.Post("/", client.createFinalFile)
		})
	})

	log.Println("Upload service starting on 0.0.0.0:8083")
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", 8083), r); err != nil {
		log.Fatal(err)
	}
}

type uploadServiceGW struct {
	// uploadpb.UnimplementedUploadServiceServer
	// client           uploadpb.UploadServiceClient
	exerciseClient   pb.ExerciseServiceClient
	postClient       pb.PostServiceClient
	submissionClient pb.SubmissionServiceClient
	attachmentClient pb.AttachmentServiceClient
}

func NewUploadService() *uploadServiceGW {
	// Specify the address and port of the gRPC gateway
	gatewayAddress := "thesis-management-backend-service:9091"

	// Create a connection to the gRPC gateway
	conn, err := grpc.Dial(gatewayAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC gateway: %v", err)
	}

	// client := uploadpb.NewUploadServiceClient(conn)
	exerciseClient := pb.NewExerciseServiceClient(conn)
	postClient := pb.NewPostServiceClient(conn)
	submissionClient := pb.NewSubmissionServiceClient(conn)
	attachmentClient := pb.NewAttachmentServiceClient(conn)

	return &uploadServiceGW{
		// client:           client,
		exerciseClient:   exerciseClient,
		postClient:       postClient,
		submissionClient: submissionClient,
		attachmentClient: attachmentClient,
	}
}
