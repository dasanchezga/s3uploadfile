package main

import (
	"aresfy_upload_ms/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func toSend() func(w http.ResponseWriter, r *http.Request) {
	s := handlers.UploadSong
	return s
}

func main() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/upload", toSend()).Methods("POST")
	r.HandleFunc("/delete", handlers.DeleteSong).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}