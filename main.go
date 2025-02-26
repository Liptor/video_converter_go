package main

import (
	"./handlers"
	"log"
	"net/http"
	"os"
)



func main() {
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/download/:filename", handlers.DownloadHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}