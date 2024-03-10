package ports

import "io"

type ImageProcessor interface {
	Process(resultWriter io.Writer, imageReader io.Reader, maxWidth, maxHeight uint) error
}
