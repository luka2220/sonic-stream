package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ApiRouter() *chi.Mux {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/files", fileRouteHandler)

	return apiRouter
}

func fileRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("file route handler"))
}