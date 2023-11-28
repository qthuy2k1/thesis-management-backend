package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
)

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

	log.Println("title", title)
	log.Println("descriptioon", description)
	log.Println("classroomID", classroomID)
	log.Println("deadline", deadline)
	log.Println("categoryID", categoryID)
	log.Println("authorID", authorID)

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

	date, err := time.Parse("2006-01-02T15:04", deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(classroomIDInt)

	res, err := u.exerciseClient.CreateExercise(ctx, &pb.CreateExerciseRequest{
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
		log.Println(err)
		log.Println(res.GetResponse().GetMessage())
		log.Println(res.GetResponse().GetStatusCode())
		http.Error(w, err.Error()+"\n"+res.GetResponse().GetMessage(), http.StatusInternalServerError)
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

func (u *uploadServiceGW) updateExercise(w http.ResponseWriter, r *http.Request) {
	log.Println("upload-service: updateExercise is called")
	// Parse the multipart form data
	err := r.ParseMultipartForm(32 << 20) // Max file size: 32MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	id, err := strconv.Atoi(chi.URLParam(r, "exerciseID"))
	if err != nil || id <= 0 {
		log.Println("id err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")
	classroomID := r.FormValue("classroomID")
	deadline := r.FormValue("deadline")
	categoryID := r.FormValue("categoryID")
	authorID := r.FormValue("authorID")

	log.Println("id", id)
	log.Println("title", title)
	log.Println("descriptioon", description)
	log.Println("classroomID", classroomID)
	log.Println("deadline", deadline)
	log.Println("categoryID", categoryID)
	log.Println("authorID", authorID)

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
		log.Println("classroomID err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		log.Println("categoryID err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02T15:04", deadline)
	if err != nil {
		log.Println("date err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := u.exerciseClient.UpdateExercise(ctx, &pb.UpdateExerciseRequest{
		Id: int64(id),
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
		log.Println(err)
		log.Println(res.GetResponse().GetMessage())
		http.Error(w, err.Error()+"\n"+res.GetResponse().GetMessage(), int(res.GetResponse().GetStatusCode()))
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
