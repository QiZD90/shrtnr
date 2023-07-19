package service

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
)

func (service *ShrtnrService) InternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (service *ShrtnrService) BadRequest(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}

func (service *ShrtnrService) Redirect(w http.ResponseWriter, url string) {
	http.Redirect(w, &http.Request{}, url, http.StatusFound)
}

func (service *ShrtnrService) ShortenLinkHandler(w http.ResponseWriter, r *http.Request) {
	link := r.PostFormValue("url")
	if _, err := url.ParseRequestURI(link); err != nil { // Validate URL
		service.BadRequest(w, "Invalid URL")
		return
	}

	shortLink, err := service.CreateLink(link)
	if err != nil {
		service.InternalServerError(w, err)
		return
	}
	fmt.Fprintf(w, `{"short_link": "%s"}`, shortLink)
}

func (service *ShrtnrService) UnshortenLinkHandler(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "link")

	link, exists, err := service.GetLink(shortLink)
	if err != nil {
		service.InternalServerError(w, err)
		return
	}

	if exists {
		fmt.Fprintf(w, `{"link": "%s"}`, link)
	} else {
		fmt.Fprintf(w, `{"error": "No such link"}`)
	}
}

func (service *ShrtnrService) LinkHandler(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "link")

	link, exists, err := service.GetLink(shortLink)
	if err != nil {
		service.InternalServerError(w, err)
		return
	}

	if exists {
		service.Redirect(w, link)
	} else {
		fmt.Fprintf(w, `{"error": "No such link"}`)
	}
}
