package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SubmissionInput struct {
	ExerciseID  string     `json:"exerciseID"`
	AuthorID    string     `json:"authorID"`
	Status      string     `json:"status"`
	Attachments []FileInfo `json:"attachments"`
}

func (u *uploadServiceGW) createSubmission(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: createSubmission is called")
	// Parse the multipart form data
	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	exerciseID := r.FormValue("exerciseID")
	if exerciseID == "" {
		http.Error(w, "Invalid exercise id", http.StatusBadRequest)
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

	submissionInput := SubmissionInput{
		ExerciseID:  exerciseID,
		AuthorID:    authorID,
		Status:      status,
		Attachments: attachments,
	}

	jsonBody, err := json.Marshal(map[string]SubmissionInput{"submission": submissionInput})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal file info: %v", err), http.StatusInternalServerError)
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://thesis-management-backend-apigw-client-service:8080/api/submit", bodyReader)
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

// func (u *uploadServiceGW) updateSubmission(w http.ResponseWriter, r *http.Request) {
// 	log.Println("upload-service: createSubmission is called")
// 	// Parse the multipart form data
// 	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
// 	if err != nil {
// 		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
// 		return
// 	}
// 	ctx := r.Context()

// 	id, err := strconv.Atoi(chi.URLParam(r, "submissionID"))
// 	if err != nil || id <= 0 {
// 		log.Println("id err", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	submissionID := r.FormValue("submissionID")
// 	authorID := r.FormValue("authorID")
// 	status := r.FormValue("status")

// 	log.Println("id", id)
// 	log.Println("submissionID", submissionID)
// 	log.Println("authorID", authorID)
// 	log.Println("status", status)

// 	// attachment
// 	var attachments []*pb.AttachmentSubmissionInput
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

// 		attachments = append(attachments, &pb.AttachmentSubmissionInput{
// 			FileURL:   fileInfo.URL,
// 			AuthorID:  authorID,
// 			Name:      fileInfo.FileName,
// 			Size:      fileInfo.Size,
// 			Type:      fileInfo.MimeType,
// 			Thumbnail: fileInfo.Thumbnail,
// 		})
// 	}

// 	submissionIDInt, err := strconv.Atoi(submissionID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	res, err := u.submissionClient.UpdateSubmission(ctx, &pb.UpdateSubmissionRequest{
// 		Id: int64(id),
// 		Submission: &pb.SubmissionInput{
// 			AuthorID:    authorID,
// 			SubmissionID:  int64(submissionIDInt),
// 			Status:      status,
// 			Attachments: attachments,
// 		},
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
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
