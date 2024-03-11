package domain

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image-resizer/internal/ports"
	"image/png"
	"io"
)

type imageProcessor struct {
}

func NewImageProcessor() ports.ImageProcessor {
	return &imageProcessor{}
}

func (i *imageProcessor) Process(resultWriter io.Writer, imageReader io.Reader, options ports.Options) error {
	// Decode image
	img, _, err := image.Decode(imageReader)
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
