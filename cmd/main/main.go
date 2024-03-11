package main

import (
	"flag"
	"fmt"
	"image-resizer/internal/domain"
	"image-resizer/internal/ports"
	_ "image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flags := parseFlags()

	flags.processFlags()

	imageProcessor := domain.NewImageProcessor()

	for _, imagePath := range flag.Args() {
		processFile(imagePath, imageProcessor, flags)
	}
}

func parseFlags() Flags {
	flags := Flags{}

	flags.maxWidth = flag.Uint("mw", 512, "width in pixels")
	flags.maxHeight = flag.Uint("mh", 512, "height in pixels")
	flags.outputDir = flag.String("o", "resized", "output directory")
	flags.help = flag.Bool("h", false, "show help")

	flag.Parse()

	if *flags.help {
		flag.Usage()
		os.Exit(0)
	}
	return flags
}

type Flags struct {
	maxWidth  *uint
	maxHeight *uint
	outputDir *string
	help      *bool
}

func (f *Flags) processFlags() {
	if len(flag.Args()) < 1 {
		fmt.Println("add path to image as argument")
		return
	}

	// creating resized directory if not exists
	err := os.Mkdir(*f.outputDir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			fmt.Println("failed to create output directory:", err)
			return
		}
	}
}

func processFile(imagePath string, imageProcessor ports.ImageProcessor, flags Flags) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			return
		}
	}(file)

	outputPath := filepath.Join(*flags.outputDir, strings.TrimSuffix(filepath.Base(imagePath), filepath.Ext(imagePath))+".png")
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("failed to create image file %s: %v\n", imagePath, err)
		return
	}
	defer func(outputFile *os.File) {
		err = outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)

	err = imageProcessor.Process(outputFile, file, ports.Options{
		MaxWidth:        *flags.maxWidth,
		MaxHeight:       *flags.maxHeight,
		SaveProportions: true,
	})
	if err != nil {
		return
	}
}
