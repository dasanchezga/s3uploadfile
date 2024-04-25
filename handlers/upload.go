package handlers

import (
	"aresfy_upload_ms/services"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// Id variable
var GlobalID string

type UploadResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func NewUploadResponse(id string, message string) *UploadResponse {
	return &UploadResponse{
		ID:      id,
		Message: message,
	}
}

func UploadSong(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("song")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// unique UUID for the file name
	id := SaveID()
	GlobalID = id
	fileName := id + ".mp3"

	// Create a temporary file to store the uploaded song
	tempFile, err := os.CreateTemp("", "uploaded_song_*.mp3")
	if err != nil {
		http.Error(w, "Unable to create temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copy the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Unable to copy file", http.StatusInternalServerError)
		return
	}
	
	// Tipo de contenido del archivo
	contentType := "audio/mpeg"
	// Upload the processed song to S3
	err = services.UploadToS3(tempFile.Name(), "aresfyuploadst",contentType, fileName)
	if err != nil {
		http.Error(w, "Unable to upload to S3", http.StatusInternalServerError)
		return
	}

	// Create the response with the generated ID and a success message
	response := UploadResponse{
		ID:      id,
		Message: "Song uploaded successfully",
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the JSON response to the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// function to unique ID
func SaveID() string {
	tempId := uuid.New()
	id := tempId.String()
	return id
}
