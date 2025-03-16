package services

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

type ImageType = string

const (
	PNG  ImageType = "png"
	JPEG ImageType = "jpeg"
	GIF  ImageType = "gif"
	BMP  ImageType = "bmp"
	WEBP ImageType = "webp"
)

type imageService struct {
	file           *multipart.FileHeader
	conversionType string
	logger         *log.Logger
}

func NewImageService(f *multipart.FileHeader, ct string) (*imageService, error) {
	logger := log.New(os.Stdout, "\u001b[34m(imageService) \u001b", 1)

	logger.Println("Image service started...")
	ifType, err := imageFileType(f)
	if err != nil {
		return nil, err
	}
	logger.Printf("Image File Type: %s\n", ifType)

	return &imageService{
		file:           f,
		conversionType: ct,
		logger:         logger,
	}, nil
}

func imageFileType(file *multipart.FileHeader) (string, error) {
	// Check if the file extension is a valid image file extension, return an error if it is not...

	n := strings.Split(file.Filename, ".")
	nImage := n[len(n)-1]

	switch nImage {
	case PNG:
		return PNG, nil
	case JPEG:
		return JPEG, nil
	case GIF:
		return GIF, nil
	case BMP:
		return BMP, nil
	case WEBP:
		return WEBP, nil
	}

	return "", errors.New(fmt.Sprintf("InvalidImageType: Got %s", nImage))
}
