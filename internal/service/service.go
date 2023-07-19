package service

import (
	"log"

	"github.com/QiZD90/shrtnr/internal/repository"
)

type ShrtnrService struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	URLPrefix  string
	Repository repository.Repository
}

func (service *ShrtnrService) CreateLink(link string) (string, error) {
	shortLink := service.GenerateShortLink()
	err := service.Repository.CreateLink(shortLink, link)
	if err != nil {
		service.ErrorLog.Print(err)
		return "", err
	}

	service.InfoLog.Printf("Shortened %s to %s", link, shortLink)
	return service.URLPrefix + shortLink, nil
}

func (service *ShrtnrService) GetLink(shortLink string) (string, bool, error) {
	link, exists, err := service.Repository.GetLink(shortLink)
	if err != nil {
		service.ErrorLog.Print(err)
		return "", false, err
	}

	service.InfoLog.Printf("Getting link %s -> %s; exists=%v", shortLink, link, exists)
	return link, exists, nil
}
