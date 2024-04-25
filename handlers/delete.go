package handlers

import (
	"aresfy_upload_ms/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteSong(w http.ResponseWriter, r *http.Request) {
    // Get the ID parameter from the request URL
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing ID parameter", http.StatusBadRequest)
        return
    }

    // Generate the S3 key based on the ID
    fileName := id + ".mp3"
    //"processed_songs/" +
    s3Key := fileName

    // Delete the file from S3
    err := services.DeleteFromS3("aresfyuploadst", s3Key)
    if err != nil {
        // Verificar si el error es específico de "archivo no encontrado"
        if err == services.ErrFileNotFound {
            // Si el archivo no se encuentra, enviar una respuesta personalizada con código de estado 404
            errorMessage := map[string]string{"error": "El archivo no existe en el bucket"}
            jsonResponse, _ := json.Marshal(errorMessage)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusNotFound)
            w.Write(jsonResponse)
            return
        }
        // Si hay otro error, devolver un error interno del servidor
        http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"message": "Song with ID %s deleted successfully from S3"}`, id)
}