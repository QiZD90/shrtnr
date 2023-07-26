package service

import (
	"errors"
	"net/url"
	"regexp"
)

var ErrUnsupportedScheme = errors.New("unsupported URI scheme")
var ErrInvalidLink = errors.New("invalid link")

var urlSchemeRegexp = regexp.MustCompile(".+://")
var httpSchemeRegexp = regexp.MustCompile("https*://")

func (s *ShrtnrService) ValidateAndFormatLink(link string) (string, error) {
	hasScheme := urlSchemeRegexp.MatchString(link)
	hasHttpScheme := httpSchemeRegexp.MatchString(link)

	if hasScheme && !hasHttpScheme { // Uh oh, some unsupported scheme
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
