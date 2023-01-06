package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const MAX_FILE_SIZE = 2048 * 2048

// handler to handle the image upload
func UploadImages(w http.ResponseWriter, r *http.Request) {
	// 32 MB is the default used by FormFile() function
	if err := r.ParseMultipartForm(MAX_FILE_SIZE); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["image"]
	var errNew string
	var http_status int
	for _, fileHeader := range files {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}
		defer file.Close()
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}
		// checking the content type
		// so we don't allow files other than images
		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
			errNew = "The provided file format is not allowed. Please upload a JPEG,JPG or PNG image"
			http_status = http.StatusBadRequest
			break
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}
		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusBadRequest
			break
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusBadRequest
			break
		}
	}
	message := "file uploaded successfully"
	messageType := "S"
	if errNew != "" {
		message = errNew
		messageType = "E"
	}
	if http_status == 0 {
		http_status = http.StatusOK
	}
	resp := map[string]interface{}{
		"messageType": messageType,
		"message":     message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http_status)
	json.NewEncoder(w).Encode(resp)
}
