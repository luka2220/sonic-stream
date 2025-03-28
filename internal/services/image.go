package services

import (
	"log/slog"
	"os"

	"github.com/luka2220/sonic-stream/internal/models/image"
)

type imageService struct {
	fileMetaData image_model.FileMetaData
	logger       *slog.Logger
}

func NewImageService(fmd image_model.FileMetaData) *imageService {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})).With("service", "image-service")

	return &imageService{
		fileMetaData: fmd,
		logger:       logger,
	}
}
