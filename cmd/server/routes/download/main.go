package download

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})).With("service", "download-api")

type DownloadRouter struct {
	Base string
	Mux  *http.ServeMux
}

func NewDownloadRouter() DownloadRouter {
	downloadMux := http.NewServeMux()
	downloadMux.HandleFunc("GET /{file}", downloadImagePath) // path parameter is the requested image file name

	return DownloadRouter{
		Base: "/download/",
		Mux:  downloadMux,
	}
}

// TODO: #5.5
// - this handler will check the ./cmd/static folder for any file matching the /download/{file} file
// - if none exists, return a client error
// - send the image image if one was found
func downloadImagePath(w http.ResponseWriter, r *http.Request) {
	logger.Info(fmt.Sprintf("%s image file requested for download", r.URL.Path))
}
