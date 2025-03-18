package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/luka2220/sonic-stream/internal/services"
)

type APIRouter struct {
	Base string
	Mux  *http.ServeMux
}

func NewAPIRoute(logger *log.Logger) *APIRouter {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("POST /image", apiImageHandler(logger))

	return &APIRouter{
		Base: "/api/",
		Mux:  apiMux,
	}
}

type serverError struct {
	Message       string
	ClientMessage string
	Status        int
}

func (e serverError) Error() string {
	return e.Message
}

type imageAPIResponse struct {
	Message string `json:"message"`
}

func createResponse(m string) ([]byte, error) {
	r := imageAPIResponse{Message: m}
	b, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("JsonMarshalError")
	}

	return b, nil
}

func validateIncomingRequest(f []*multipart.FileHeader, c string) error {
	// Currently only accepting 1 image file
	if len(f) != 0 {
		return &serverError{
			Message:       fmt.Sprintf("400 POST <- client supplied invalid image file length: %d\n", len(f)),
			ClientMessage: fmt.Sprintf("Invalid image file length: %d, need 1", len(f)),
			Status:        400,
		}
	}

	if c == "" {
		return &serverError{
			Message:       fmt.Sprintf("400 POST <- client supplied no convert file type"),
			ClientMessage: fmt.Sprintf("Invalid post schema, expecting string value 'convert'"),
			Status:        400,
		}
	}

	return nil
}

func internalServerError(w http.ResponseWriter, m string, l *log.Logger) {
	l.Println(m)
	w.WriteHeader(500)
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Server Error"))
}

// NOTE: This handler is currently too big and hard to tell what happens where:
// - Parsing multipart form data
// - Validating input
// - Handeling diffrent error casses (400, 500)
// - Creating JSON response
// Accepts an image file of up to 150kb
func apiImageHandler(logger *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json'")

		if err := r.ParseMultipartForm(250000); err != nil {
			e := fmt.Sprintf("Error parsing form data from client: %s\n", err.Error())
			internalServerError(w, e, logger)
			return
		}

		// TODO: Should also check for valid image file extensions and image file size when validing the incoming request
		imageFile := r.MultipartForm.File["file"]
		convertType := r.PostForm.Get("convert")
		if err := validateIncomingRequest(imageFile, convertType); err != nil {
			ok := err.(serverError)
			if ok.Status == 500 {
				internalServerError(w, err.Error(), logger)
				return
			}

			logger.Println(ok.Error())
			b, err := createResponse(ok.ClientMessage)
			if err != nil {
				internalServerError(w, err.Error(), logger)
				return
			}
			w.WriteHeader(ok.Status)
			w.Write(b)
			return
		}

		i, err := services.NewImageService(imageFile[0], r.PostForm.Get("convert"))
		if err != nil {
			ok := err.(services.HttpError)
			switch ok.Status {
			case 500:
				e := fmt.Sprintf("\u001b[31m[ERROR]: %s\u001b\n", ok.Error())
				internalServerError(w, e, logger)
				return
			case 400:
				logger.Println(ok.Error())
				w.Header().Add("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(400)
				w.Write([]byte(ok.Error()))
				return
			}

			e := fmt.Sprintf("unexpected error type: %s\n", err.Error())
			internalServerError(w, e, logger)
			return
		}

		msg := fmt.Sprintf("%s -> %s", i.BaseExtension, i.ConvertType)
		bytes, err := createResponse(msg)
		if err != nil {
			internalServerError(w, err.Error(), logger)
			return
		}

		w.WriteHeader(200)
		w.Write(bytes)
	}
}
