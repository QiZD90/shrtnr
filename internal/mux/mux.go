package mux

import (
	"net/http"

	"github.com/QiZD90/shrtnr/internal/service"
	"github.com/go-chi/chi"
)

func Get(service *service.ShrtnrService) http.Handler {
	mux := chi.NewRouter()

	mux.Post("/shorten", service.ShortenLinkHandler)
	mux.Get("/unshorten/{link}", service.UnshortenLinkHandler)
	mux.Get("/{link}", service.LinkHandler)

	return mux
}
