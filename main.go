package main

import (
	"video_converter_go/handlers"
	"log"
	"net/http"
)



func main() {
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download/:filename", handlers.DownloadHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}