package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PostInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ClassroomID string     `json:"classroomID"`
	CategoryID  string     `json:"categoryID"`
	AuthorID    string     `json:"authorID"`
	Attachments []FileInfo `json:"attachments"`
}

func (u *uploadServiceGW) createPost(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: createPost is called")
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

	postInput := PostInput{
		Title:       title,
		Description: description,
		ClassroomID: classroomID,
		CategoryID:  categoryID,
		AuthorID:    authorID,
		Attachments: attachments,
	}

	jsonBody, err := json.Marshal(map[string]PostInput{"post": postInput})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal file info: %v", err), http.StatusInternalServerError)
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://thesis-management-backend-apigw-client-service:8080/api/post", bodyReader)
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

func (u *uploadServiceGW) updatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: updatePost is called")
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

	status := r.FormValue("status")

	postID := chi.URLParam(r, "postID")

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

	postInput := PostInput{
		Title:       title,
		Description: description,
		ClassroomID: classroomID,
		CategoryID:  categoryID,
		AuthorID:    authorID,
		Attachments: attachments,
	}

	jsonBody, err := json.Marshal(map[string]PostInput{"post": postInput})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal file info: %v", err), http.StatusInternalServerError)
	}

	url := "http://thesis-management-backend-apigw-client-service:8080/api/post/" + postID

	log.Println(url)

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
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
