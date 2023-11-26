package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
)

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
	description := r.FormValue("description")
	classroomID := r.FormValue("classroomID")
	categoryID := r.FormValue("categoryID")
	authorID := r.FormValue("authorID")

	// attachment
	var attachments []*pb.AttachmentPostInput
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

		attachments = append(attachments, &pb.AttachmentPostInput{
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

	res, err := u.postClient.CreatePost(ctx, &pb.CreatePostRequest{
		Post: &pb.PostInput{
			Title:       title,
			Description: description,
			ClassroomID: int64(classroomIDInt),
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
