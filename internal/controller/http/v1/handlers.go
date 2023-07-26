package v1

import (
	"encoding/json"
	"net/http"

	"github.com/QiZD90/shrtnr/internal/service"
	"github.com/go-chi/chi"
)

type shrtnrRoutes struct {
	s *service.ShrtnrService
}

// Used only for reporting errors, that occured while parsing json response
// If you want to report an error, do it through JsonError
func (routes *shrtnrRoutes) InternalServerError(w http.ResponseWriter, err error) {
	routes.s.ErrorLog.Print(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (routes *shrtnrRoutes) Redirect(w http.ResponseWriter, url string) {
	http.Redirect(w, &http.Request{}, url, http.StatusFound)
}

func (routes *shrtnrRoutes) RespondWithJson(w http.ResponseWriter, statusCode int, j JsonResponse) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")

	b, err := j.ToBytes()
	if err != nil {
		routes.InternalServerError(w, err)
		return
	}

	w.Write(b)
}

func (routes *shrtnrRoutes) ShortenLinkHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request
	var j JsonShortenRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&j); err != nil {
		routes.RespondWithJson(w, http.StatusBadRequest, &JsonError{"Error while parsing request JSON"})
		routes.s.ErrorLog.Print(err)
		return
	}

	// Validate URL
	link, err := routes.s.ValidateAndFormatLink(j.Link)
	if err != nil {
		routes.RespondWithJson(w, http.StatusBadRequest, &JsonError{"Invalid URL"})
		return
	}

	shortLink, err := routes.s.CreateLink(link)
	if err != nil {
		routes.RespondWithJson(w, http.StatusInternalServerError, &JsonError{"Error while shortening link"})
		routes.s.ErrorLog.Print(err)
		return
	}

	routes.RespondWithJson(w, http.StatusOK, &JsonShortenResponse{shortLink})
}

func (routes *shrtnrRoutes) UnshortenLinkHandler(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "link")

	fullLink, exists, err := routes.s.GetLink(shortLink)
	if err != nil {
		routes.RespondWithJson(w, http.StatusInternalServerError, &JsonError{"Error while unshortening link"})
		routes.s.ErrorLog.Print(err)
		return
	}

	if exists {
		routes.RespondWithJson(w, http.StatusOK, &JsonUnshortenResponse{fullLink})
	} else {
		routes.RespondWithJson(w, http.StatusNotFound, &JsonError{"No such link"})
	}
}

func (routes *shrtnrRoutes) LinkHandler(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "link")

	fullLink, exists, err := routes.s.GetLink(shortLink)
	if err != nil {
		routes.RespondWithJson(w, http.StatusInternalServerError, &JsonError{"Error while unshortening link"})
		routes.s.ErrorLog.Print(err)
		return
	}

	if exists {
		routes.Redirect(w, fullLink)
	} else {
		routes.RespondWithJson(w, http.StatusNotFound, &JsonError{"No such link"})
	}
}
