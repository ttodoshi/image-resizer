package api

import (
	"image-resizer/internal/domain"
	"image-resizer/internal/ports"
	"mime/multipart"
	"net/http"
	"strconv"
)

var imageProcessor = domain.NewImageProcessor()

func ResizeImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, err, closeFunc := parseMultipartFile(r)
	if err != nil {
		return
	}
	defer closeFunc()

	width, err := strconv.Atoi(r.URL.Query().Get("width"))
	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = imageProcessor.Process(w, file, ports.Options{
		MaxWidth:        uint(width),
		MaxHeight:       uint(height),
		SaveProportions: r.URL.Query().Get("save-proportions") != "false",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseMultipartFile(r *http.Request) (multipart.File, error, func()) {
	err := r.ParseMultipartForm(10 << 20) // max size 10MB
	if err != nil {
		return nil, err, nil
	}

	file, _, err := r.FormFile("file")
	return file, err, func() {
		err = file.Close()
		if err != nil {
			return
		}
	}
}
