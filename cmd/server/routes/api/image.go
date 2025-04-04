package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	image_model "github.com/luka2220/sonic-stream/internal/models/image"
	"github.com/luka2220/sonic-stream/internal/services"
)

var validImageTypes = map[string]string{
	"PNG":  "png",
	"JPEG": "jpeg",
	"GIF":  "gif",
	"BMP":  "bmp",
	"WEBP": "webp",
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
})).With("service", "image-api")

type imageAPIResponse struct {
	DownloadUrl string `json:"download_url"`
}

func createResponse(m string) ([]byte, error) {
	url := fmt.Sprintf("http://localhost:8080/download/%s", m)
	r := imageAPIResponse{DownloadUrl: url}
	b, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("JsonMarshalError")
	}

	return b, nil
}

// Validate the incoming files extension
func validateImageFileExtensions(f *multipart.FileHeader, c string) (image_model.FileMetaData, error) {
	validImageType := false
	validConvertType := false

	imageModel := image_model.FileMetaData{}

	n := strings.Split(f.Filename, ".")
	imgType := n[len(n)-1]

	for _, image := range validImageTypes {
		if imgType == image {
			validImageType = true
			imageModel.BaseExtention = image
			imageModel.Base = f
			continue
		}
		if c == image {
			validConvertType = true
			imageModel.ConvertExtension = c
			continue
		}
	}

	if validImageType && validConvertType {
		return imageModel, nil
	}

	errMsg := fmt.Sprintf("400 POST <- invalid file extensions, got=(%s, %s)", imgType, c)
	clientResponseMsg := fmt.Sprintf("Invalid file extension type: %s, %s (check docs for valid image extensions)", imgType, c)

	return image_model.FileMetaData{}, ServerError{Message: errMsg, ClientMessage: clientResponseMsg, Status: 400}
}

// Validate the incoming request has both file field and convert field
func validateIncomingRequest(f []*multipart.FileHeader, c string) (image_model.FileMetaData, error) {
	if len(f) != 1 {
		return image_model.FileMetaData{}, ServerError{
			Message:       fmt.Sprintf("400 POST <- client supplied invalid image file length: %d\n", len(f)),
			ClientMessage: fmt.Sprintf("Invalid image file length: %d, need 1", len(f)),
			Status:        400,
		}
	}

	if c == "" {
		return image_model.FileMetaData{}, ServerError{
			Message:       "400 POST <- client supplied no convert file type",
			ClientMessage: "Invalid post schema, expecting string value 'convert'",
			Status:        400,
		}
	}

	r, err := validateImageFileExtensions(f[0], c)
	if err != nil {
		return image_model.FileMetaData{}, err
	}

	return r, nil
}

func apiImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Header().Add("Content-Type", "application/json'")

	if err := r.ParseMultipartForm(250000); err != nil {
		InternalServerError(ServerError{
			Message:       err.Error(),
			W:             w,
			Status:        500,
			L:             logger,
			ClientMessage: "Server Error",
		})
		return
	}

	imageFile := r.MultipartForm.File["file"]
	convertType := r.PostForm.Get("convert")

	imageMetaData, err := validateIncomingRequest(imageFile, convertType)
	if err != nil {
		e := err.(ServerError)
		e.L = logger
		e.W = w
		ClientError(e)
		return
	}

	is := services.NewImageService(imageMetaData)
	convertedFileName, err := is.GetConvertedImage()
	if err != nil {
		InternalServerError(ServerError{
			err.Error(),
			"something went wrong on our end",
			500,
			w,
			logger,
		})
		return
	}

	bytes, err := createResponse(convertedFileName)
	if err != nil {
		InternalServerError(ServerError{err.Error(), "something went wrong on our end", 500, w, logger})
		return
	}

	w.WriteHeader(200)
	w.Write(bytes)
	logger.Info(fmt.Sprintf("200 %s created under ./cmd/static", convertedFileName))
}
