package v1

import (
	"net/http"

	"github.com/QiZD90/shrtnr/internal/service"
	"github.com/go-chi/chi"
)

func NewMux(s *service.ShrtnrService) http.Handler {
	mux := chi.NewRouter()
	routes := &shrtnrRoutes{
		s: s,
	}

	mux.Post("/shorten", routes.ShortenLinkHandler)
	mux.Get("/unshorten/{link:[a-z]+-[a-z]+-[a-z]+}", routes.UnshortenLinkHandler)
	mux.Get("/{link:[a-z]+-[a-z]+-[a-z]+}", routes.LinkHandler)

	return mux
}
