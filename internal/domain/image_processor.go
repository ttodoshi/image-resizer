package domain

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image-cropper/internal/ports"
	"image/png"
	"io"
)

type imageProcessor struct {
}

func NewImageProcessor() ports.ImageProcessor {
	return &imageProcessor{}
}

func (i *imageProcessor) Process(resultWriter io.Writer, imageReader io.Reader, maxWidth, maxHeight uint) error {
	// Decode image
	img, _, err := image.Decode(imageReader)
	if err != nil {
		fmt.Printf("error decoding image: %v\n", err)
		return err
	}

	// Determine new dimensions while preserving aspect ratio
	var newWidth, newHeight uint
	width := uint(img.Bounds().Dx())
	height := uint(img.Bounds().Dy())

	if width > height {
		newWidth = maxWidth
		newHeight = maxHeight * height / width
	} else {
		newHeight = maxHeight
		newWidth = maxWidth * width / height
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
