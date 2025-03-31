package api

import (
	"net/http"
)

type APIRouter struct {
	Base string
	Mux  *http.ServeMux
}

func NewAPIRoute() *APIRouter {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("POST /image", apiImageHandler)

	return &APIRouter{
		Base: "/api/",
		Mux:  apiMux,
	}
}
