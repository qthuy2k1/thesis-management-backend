package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/genproto/googleapis/type/datetime"
)

type ExerciseInput struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	ClassroomID string            `json:"classroomID"`
	CategoryID  string            `json:"categoryID"`
	AuthorID    string            `json:"authorID"`
	Deadline    datetime.DateTime `json:"deadline"`
	Attachments []FileInfo        `json:"attachments"`
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
	if title == "" {
		http.Error(w, "Missing title", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	if description == "" {
		http.Error(w, "Missing description", http.StatusBadRequest)
		return
	}

	classroomID := r.FormValue("classroomID")
	if classroomID == "" {
		http.Error(w, "Invalid classroom id", http.StatusBadRequest)
		return
	}

	categoryID := r.FormValue("categoryID")
	if categoryID == "" {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	authorID := r.FormValue("authorID")
	if authorID == "" {
		http.Error(w, "Invalid author id", http.StatusBadRequest)
		return
	}

	deadline := r.FormValue("deadline")
	if deadline == "" {
		http.Error(w, "Invalid deadline", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02T15:04", deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := r.FormValue("status")

	// attachment
	var attachments []FileInfo
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

		attachments = append(attachments, FileInfo{
			Name:      driveFile.Name,
			Thumbnail: driveFile.ThumbnailLink,
			Size:      driveFile.Size,
			Type:      driveFile.MimeType,
			FileURL:   driveFile.WebViewLink,
			AuthorID:  authorID,
			Status:    status,
		})
	}

	exerciseInput := ExerciseInput{
		Title:       title,
		Description: description,
		ClassroomID: classroomID,
		CategoryID:  categoryID,
		AuthorID:    authorID,
		Attachments: attachments,
		Deadline: datetime.DateTime{
			Year:    int32(date.Year()),
			Month:   int32(date.Month()),
			Day:     int32(date.Day()),
			Hours:   int32(date.Hour()),
			Minutes: int32(date.Minute()),
		},
	}

	jsonBody, err := json.Marshal(map[string]ExerciseInput{"exercise": exerciseInput})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal file info: %v", err), http.StatusInternalServerError)
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://thesis-management-backend-apigw-client-service:8080/api/exercise", bodyReader)
	if err != nil {
		http.Error(w, fmt.Sprintf("make request to apigw client failed: %v", err), http.StatusBadGateway)
		return
	}

	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		http.Error(w, "Authorization is required", http.StatusMethodNotAllowed)
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authorization)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("make request to apigw client failed: %v", err), http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var respBody Response
	if err := json.Unmarshal(body, &respBody); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal response body: %v", err), http.StatusInternalServerError)
	}

	w.Write(body)
}

// func (u *uploadServiceGW) updateExercise(w http.ResponseWriter, r *http.Request) {
// 	log.Println("upload-service: updateExercise is called")
// 	// Parse the multipart form data
// 	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
// 	if err != nil {
// 		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
// 		return
// 	}
// 	ctx := r.Context()

// 	id, err := strconv.Atoi(chi.URLParam(r, "exerciseID"))
// 	if err != nil || id <= 0 {
// 		log.Println("id err", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	title := r.FormValue("title")
// 	description := r.FormValue("description")
// 	classroomID := r.FormValue("classroomID")
// 	deadline := r.FormValue("deadline")
// 	categoryID := r.FormValue("categoryID")
// 	authorID := r.FormValue("authorID")

// 	log.Println("id", id)
// 	log.Println("title", title)
// 	log.Println("descriptioon", description)
// 	log.Println("classroomID", classroomID)
// 	log.Println("deadline", deadline)
// 	log.Println("categoryID", categoryID)
// 	log.Println("authorID", authorID)

// 	// attachment
// 	var attachments []*pb.AttachmentExerciseInput
// 	fhs := r.MultipartForm.File["attachments"]
// 	for _, fileHeader := range fhs {
// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer file.Close()

// 		driveFile, err := uploadFileToDrive(ctx, file, fileHeader)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		fileInfo := FileInfo{
// 			FileName:  driveFile.Name,
// 			Thumbnail: driveFile.ThumbnailLink,
// 			Size:      driveFile.Size,
// 			MimeType:  driveFile.MimeType,
// 			URL:       driveFile.WebViewLink,
// 		}

// 		attachments = append(attachments, &pb.AttachmentExerciseInput{
// 			FileURL:   fileInfo.URL,
// 			AuthorID:  authorID,
// 			Name:      fileInfo.FileName,
// 			Size:      fileInfo.Size,
// 			Type:      fileInfo.MimeType,
// 			Thumbnail: fileInfo.Thumbnail,
// 		})
// 	}

// 	classroomIDInt, err := strconv.Atoi(classroomID)
// 	if err != nil {
// 		log.Println("classroomID err", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	categoryIDInt, err := strconv.Atoi(categoryID)
// 	if err != nil {
// 		log.Println("categoryID err", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	date, err := time.Parse("2006-01-02T15:04", deadline)
// 	if err != nil {
// 		log.Println("date err", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	res, err := u.exerciseClient.UpdateExercise(ctx, &pb.UpdateExerciseRequest{
// 		Id: int64(id),
// 		Exercise: &pb.ExerciseInput{
// 			Title:       title,
// 			Description: description,
// 			ClassroomID: int64(classroomIDInt),
// 			Deadline: &datetime.DateTime{
// 				Year:    int32(date.Year()),
// 				Month:   int32(date.Month()),
// 				Day:     int32(date.Day()),
// 				Hours:   int32(date.Hour()),
// 				Minutes: int32(date.Minute()),
// 				Seconds: int32(date.Second()),
// 			},
// 			CategoryID:  int64(categoryIDInt),
// 			AuthorID:    authorID,
// 			Attachments: attachments,
// 		},
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		log.Println(res.GetResponse().GetMessage())
// 		http.Error(w, err.Error()+"\n"+res.GetResponse().GetMessage(), int(res.GetResponse().GetStatusCode()))
// 		return
// 	}

// 	response := map[string]interface{}{
// 		"code":    res.Response.StatusCode,
// 		"message": res.Response.Message,
// 	}
// 	jsonBytes, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(int(res.Response.StatusCode))
// 	w.Write(jsonBytes)
// }
