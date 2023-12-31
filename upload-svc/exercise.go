package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

func (u *uploadServiceGW) updateExercise(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: updateExercise is called")
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

	exerciseID := chi.URLParam(r, "exerciseID")

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
	req, err := http.NewRequest(http.MethodPut, "http://thesis-management-backend-apigw-client-service:8080/api/exercise/"+exerciseID, bodyReader)
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
