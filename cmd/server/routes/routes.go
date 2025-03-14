package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type APIRouter struct {
	Base string
	Mux *http.ServeMux
}

type imageAPIResponse struct {
	Message string `json:"message"`;
}

func NewAPIRoute(logger *log.Logger) *APIRouter {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("POST /image", apiImageHandler(logger))

	return &APIRouter{
		Base: "/api/",
		Mux: apiMux,
	}
}

// Accepts an image file of up to 500kb
func apiImageHandler(logger *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(500000); err != nil {
			logger.Fatalf("Error parsing form data from client: %s", err.Error())
			w.WriteHeader(500)
			w.Write([]byte("Server Error"))
			return
		}

		logger.Println(r.PostForm)
		image_file := r.MultipartForm.File["file"]

		if len(image_file) == 0 {
			w.WriteHeader(400)

			logger.Println("POST -> /api/image: client missing file key/value in form data")

			response := &imageAPIResponse{}
			response.Message = "Missing image file in POST form-data. i.e 'file': 'image_file.png'"
			byteResponse, err := json.Marshal(response)
			
			if err != nil {
				logger.Printf("\u001b[31m[ERROR]: %s\u001b\n", err.Error())
				
				w.WriteHeader(500)
				w.Write([]byte("Server Error"))
				return 
			}

			w.Write(byteResponse)
		}

		f := image_file[0]
		logger.Printf("\u001b[32mFile Name:\u001b[0m %s\n", f.Filename)
		logger.Printf("\u001b[36mFile Size:\u001b[0m %d bytes\n", f.Size)

		f_reader, err := f.Open()
		if err != nil {
			logger.Printf("\u001b[31m[ERROR]: %s\u001b\n", err.Error())
				
			w.WriteHeader(500)
			w.Write([]byte("Server Error"))
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
				logger.Printf("\u001b[31m[ERROR]: %s\u001b\n", err.Error())
				
				w.WriteHeader(500)
				w.Write([]byte("Server Error"))
				return 
			}

			logger.Printf("bytes read: %v\n", b)
		}
	}
}
