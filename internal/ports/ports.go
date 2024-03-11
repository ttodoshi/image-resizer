package ports

import "io"

type Options struct {
	MaxWidth        uint
	MaxHeight       uint
	SaveProportions bool
}

type ImageProcessor interface {
	Process(resultWriter io.Writer, imageReader io.Reader, options Options) error
}
