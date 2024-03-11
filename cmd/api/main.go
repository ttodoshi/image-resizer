package main

import (
	"image-resizer/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.ResizeImage)

	log.Println("Starting server on port :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
