package service

import (
	"log"
	"net/url"

	"github.com/QiZD90/shrtnr/internal/repository"
)

type ShrtnrService struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	URLPrefix  string
	Repository repository.Repository
}

func (s *ShrtnrService) CreateLink(link string) (string, error) {
	shortLink := s.GenerateShortLink()
	err := s.Repository.CreateLink(shortLink, link)
	if err != nil {
		s.ErrorLog.Print(err)
		return "", err
	}

	s.InfoLog.Printf("Shortened %s to %s", link, shortLink)
	return url.JoinPath(s.URLPrefix, shortLink)
}

func (s *ShrtnrService) GetLink(shortLink string) (string, bool, error) {
	link, exists, err := s.Repository.GetLink(shortLink)
	if err != nil {
		s.ErrorLog.Print(err)
		return "", false, err
	}

	s.InfoLog.Printf("Getting link %s -> %s; exists=%v", shortLink, link, exists)
	return link, exists, nil
}
