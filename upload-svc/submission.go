package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
)

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
	authorID := r.FormValue("authorID")
	submissionDate := r.FormValue("submissionDate")
	status := r.FormValue("status")

	// attachment
	var attachments []*pb.AttachmentSubmissionInput
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

		attachments = append(attachments, &pb.AttachmentSubmissionInput{
			FileURL:   fileInfo.URL,
			AuthorID:  authorID,
			Name:      fileInfo.FileName,
			Size:      fileInfo.Size,
			Type:      fileInfo.MimeType,
			Thumbnail: fileInfo.Thumbnail,
		})
	}

	exerciseIDInt, err := strconv.Atoi(exerciseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02 15:04:05", submissionDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.submissionClient.CreateSubmission(ctx, &pb.CreateSubmissionRequest{
		Submission: &pb.SubmissionInput{
			UserID:     authorID,
			ExerciseID: int64(exerciseIDInt),
			SubmissionDate: &datetime.DateTime{
				Year:    int32(date.Year()),
				Month:   int32(date.Month()),
				Day:     int32(date.Day()),
				Hours:   int32(date.Hour()),
				Minutes: int32(date.Minute()),
				Seconds: int32(date.Second()),
			},
			Status:      status,
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
