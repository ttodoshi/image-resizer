package decoder

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

func DecodeImage(file multipart.File) (image.Image, error) {
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}

	if _, err = file.Seek(0, 0); err != nil {
		return nil, err
	}

	var img image.Image
	switch format {
	case "jpg":
	case "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return nil, err
	}
	return img, nil
}
