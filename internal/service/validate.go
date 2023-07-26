package service

import (
	"errors"
	"net/url"
	"strings"
)

var ErrUnsupportedScheme = errors.New("unsupported URI scheme")
var ErrInvalidLink = errors.New("invalid link")

func (s *ShrtnrService) ValidateAndFormatLink(link string) (string, error) {
	hasScheme := strings.Contains(link, "://")

	if hasScheme && !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		return "", ErrUnsupportedScheme
	} else if !hasScheme { // Assume https:// if we don't have a scheme
		link = "https://" + link
	}

	sanitizedUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return "", ErrInvalidLink
	}

	return sanitizedUrl.String(), nil
}
