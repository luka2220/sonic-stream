package services

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

var validImageTypes = map[string]string{
	"PNG":  "png",
	"JPEG": "jpeg",
	"GIF":  "gif",
	"BMP":  "bmp",
	"WEBP": "webp",
}

type imageService struct {
	file          *multipart.FileHeader
	BaseExtension string
	ConvertType   string
	logger        *log.Logger
}

// Custom error messages including http status
type HttpError struct {
	Status  int
	Message string
}

func (h HttpError) Error() string {
	return h.Message
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
		file:          f,
		ConvertType:   ct,
		logger:        logger,
		BaseExtension: ifType,
	}, nil
}

// Checks if the image file is a valid type the service accepts
func imageFileType(file *multipart.FileHeader) (string, error) {
	n := strings.Split(file.Filename, ".")
	nImage := n[len(n)-1]

	for _, image := range validImageTypes {
		if nImage == image {
			return image, nil
		}
	}

	e := fmt.Sprintf("InvalidImageType: Got %s", nImage)

	return "", HttpError{Message: e, Status: 400}
}
