package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/luka2220/sonic-stream/internal/services"
)

type APIRouter struct {
	Base string
	Mux  *http.ServeMux
}

type imageAPIResponse struct {
	Message string `json:"message"`
}

func NewAPIRoute(logger *log.Logger) *APIRouter {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("POST /image", apiImageHandler(logger))

	return &APIRouter{
		Base: "/api/",
		Mux:  apiMux,
	}
}

func internalServerError(w http.ResponseWriter, m string, l *log.Logger) {
	l.Println(m)
	w.WriteHeader(500)
	w.Write([]byte("Server Error"))
}

// Accepts an image file of up to 500kb
func apiImageHandler(logger *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set http response type
		w.Header().Add("Content-Type", "application/json'")

		if err := r.ParseMultipartForm(500000); err != nil {
			internalServerError(w, fmt.Sprintf("Error parsing form data from client: %s\n", err.Error()), logger)
			return
		}

		imageFile := r.MultipartForm.File["file"]
		if convertType := r.PostForm.Get("convert"); convertType == "" {
			logger.Println("No convert type sent in post-form from client")

			r := &imageAPIResponse{
				Message: "Invalid post-form, missing convert key/value",
			}

			rBytes, err := json.Marshal(r)
			if err != nil {
				internalServerError(w, fmt.Sprintf("Error serializing struct to json: %s\n", err.Error()), logger)
				return
			}

			w.WriteHeader(400)
			w.Write(rBytes)
			return
		}

		if len(imageFile) == 0 {
			logger.Println("POST -> /api/image: client missing file key/value in form data")

			response := &imageAPIResponse{}
			response.Message = "Missing image file in POST form-data. i.e 'file': 'image_file.png'"
			byteResponse, err := json.Marshal(response)

			if err != nil {
				internalServerError(w, fmt.Sprintf("\u001b[31m[ERROR]: %s\u001b\n", err.Error()), logger)
				return
			}

			w.WriteHeader(400)
			w.Write(byteResponse)
			return
		}

		// Extract functionality to ImageService???

		f := imageFile[0]
		logger.Printf("\u001b[32mFile Name:\u001b[0m %s\n", f.Filename)
		logger.Printf("\u001b[36mFile Size:\u001b[0m %d bytes\n", f.Size)

		_, err := services.NewImageService(f, r.PostForm.Get("convert"))
		if err != nil {
			// For now leave as internal server error
			internalServerError(w, err.Error(), logger)
			return
		}

		f_reader, err := f.Open()
		if err != nil {
			internalServerError(w, fmt.Sprintf("\u001b[31m[ERROR]: %s\u001b\n", err.Error()), logger)
			return
		}

		for {
			// Read image data by the byte
			b := make([]byte, 8)
			_, err := f_reader.Read(b)
			if err == io.EOF {
				f_reader.Close()
				break
			}

			if err != nil {
				internalServerError(w, fmt.Sprintf("\u001b[31m[ERROR]: %s\u001b\n", err.Error()), logger)
				return
			}

			// logger.Printf("bytes read: %v\n", b)
		}
	}
}
