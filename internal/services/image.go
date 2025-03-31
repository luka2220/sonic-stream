package services

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log/slog"
	"os"

	"github.com/google/uuid"

	image_model "github.com/luka2220/sonic-stream/internal/models/image"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})).With("service", "image-service")

type imageService struct {
	image_model.FileMetaData
}

func NewImageService(fmd image_model.FileMetaData) *imageService {
	return &imageService{fmd}
}

// TODO:
// - Convert an image file from one type to another
// - Store the converted file in ./cmd/static
// - Create the download url endpoint to send back to the client
func (is imageService) GetConvertedImage() (string, error) {
	switch is.BaseExtention {
	case "png":
		return is.convertPNG()
	case "jpeg":
		return "", errors.New("jpeg conversion not implemented")
	case "gif":
		return "", errors.New("gif conversion not implemented")
	case "bmp":
		return "", errors.New("bmp conversion not implemented")
	case "webp":
		return "", errors.New("webp conversion not implemented")
	}

	return "", nil
}

func (is imageService) convertPNG() (string, error) {
	file, err := is.Base.Open()
	if err != nil {
		logger.Error("Error opening multipart file header")
		return "", err
	}

	defer file.Close()

	image, err := png.Decode(file)
	if err != nil {
		logger.Error("Error decoding png image")
		return "", err
	}

	switch is.ConvertExtension {
	case "jpeg":
		path, err := encodeJPEG(image)
		if err != nil {
			return "", err
		}
		return path, nil
	case "gif":
		return "", errors.New("png to gif not yet implemented")
	case "bmp":
		return "", errors.New("png bmp not implemented")
	case "webp":
		return "", errors.New("png to webp gif not implemented")
	}

	return "", nil
}

// Encode the base image.Image into a .jpeg file format
// Creates a file stored at ./cmd/static that hold the contents of the jpeg image
// Returns the jpeg file name created
func encodeJPEG(i image.Image) (string, error) {
	fid := uuid.NewString()
	fName := fmt.Sprintf("%s.jpeg", fid)
	path := fmt.Sprintf("./cmd/static/%s", fName)
	f, err := os.Create(path)
	if err != nil {
		logger.Error("Error creating file path for the generated jpeg")
		return "", err
	}

	err = jpeg.Encode(f, i, nil)
	if err != nil {
		logger.Error("Error encoding jpeg file")
		return "", err
	}

	return fName, nil
}
