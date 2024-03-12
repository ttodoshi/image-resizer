package handler

import (
	"fmt"
	"github.com/nfnt/resize"
	"image-resizer/pkg/image/decoder"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

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

	err = ProcessImage(w, file, Options{
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

type Options struct {
	MaxWidth        uint
	MaxHeight       uint
	SaveProportions bool
}

func ProcessImage(resultWriter io.Writer, file multipart.File, options Options) error {
	// Decode image
	img, err := decoder.DecodeImage(file)
	if err != nil {
		fmt.Printf("error decoding image: %v\n", err)
		return err
	}

	var newWidth, newHeight uint
	if options.SaveProportions {
		width := uint(img.Bounds().Dx())
		height := uint(img.Bounds().Dy())

		if width > height {
			newWidth = options.MaxWidth
			newHeight = options.MaxHeight * height / width
		} else {
			newHeight = options.MaxHeight
			newWidth = options.MaxWidth * width / height
		}
	} else {
		newWidth = options.MaxWidth
		newHeight = options.MaxHeight
	}

	// Resize the image with calculated dimensions
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	err = png.Encode(resultWriter, resizedImg)
	if err != nil {
		fmt.Printf("error writing image: %v\n", err)
		return err
	}

	return nil
}
