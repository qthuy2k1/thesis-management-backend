package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func (u *uploadServiceGW) createFinalFile(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: createFinalFile is called")
	// Parse the multipart form data
	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	authorID := r.FormValue("authorID")

	if strings.TrimSpace(authorID) == "" {
		http.Error(w, "missing authorID", http.StatusBadRequest)
		return
	}

	status := r.FormValue("status")

	// attachment

	fhs := r.MultipartForm.File["attachments"]
	if len(fhs) == 0 {
		http.Error(w, "missing file to upload", http.StatusBadRequest)
		return
	}

	file, err := fhs[0].Open()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	driveFile, err := uploadFileToDrive(ctx, file, fhs[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileInfo := FileInfo{
		Name:      driveFile.Name,
		Thumbnail: driveFile.ThumbnailLink,
		Size:      driveFile.Size,
		Type:      driveFile.MimeType,
		FileURL:   driveFile.WebViewLink,
		AuthorID:  authorID,
		Status:    status,
	}

	jsonBody, err := json.Marshal(map[string]FileInfo{"attachment": fileInfo})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal file info: %v", err), http.StatusInternalServerError)
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://thesis-management-backend-apigw-client-service:8080/api/attachment", bodyReader)
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
	req.Header.Set("Authorization", "Bearer "+authorization)

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

// func (u *uploadServiceGW) updateFinalFile(w http.ResponseWriter, r *http.Request) {
// 	log.Println("upload-service: updateFinalFile is called")
// 	// Parse the multipart form data
// 	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
// 	if err != nil {
// 		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
// 		return
// 	}
// 	ctx := r.Context()

// 	id, err := strconv.Atoi(chi.URLParam(r, "finalFileID"))
// 	if err != nil || id <= 0 {
// 		log.Println("id err", err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	title := r.FormValue("title")
// 	description := r.FormValue("description")
// 	classroomID := r.FormValue("classroomID")
// 	categoryID := r.FormValue("categoryID")
// 	authorID := r.FormValue("authorID")

// 	log.Println("id", id)
// 	log.Println("title", title)
// 	log.Println("descriptioon", description)
// 	log.Println("classroomID", classroomID)
// 	log.Println("categoryID", categoryID)
// 	log.Println("authorID", authorID)

// 	// attachment
// 	var attachments []*pb.AttachmentFinalFileInput
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

// 		attachments = append(attachments, &pb.AttachmentFinalFileInput{
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

// 	res, err := u.finalFileClient.UpdateFinalFile(ctx, &pb.UpdateFinalFileRequest{
// 		Id: int64(id),
// 		FinalFile: &pb.FinalFileInput{
// 			Title:       title,
// 			Description: description,
// 			ClassroomID: int64(classroomIDInt),
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
