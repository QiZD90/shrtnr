package v1

import "encoding/json"

type JsonResponse interface {
	ToBytes() ([]byte, error)
}

type JsonError struct {
	Message string `json:"error"`
}

func (j *JsonError) ToBytes() ([]byte, error) {
	return json.Marshal(j)
}

type JsonShortenResponse struct {
	ShortLink string `json:"short_link"`
}

func (j *JsonShortenResponse) ToBytes() ([]byte, error) {
	return json.Marshal(j)
}

type JsonUnshortenResponse struct {
	FullLink string `json:"full_link"`
}

func (j *JsonUnshortenResponse) ToBytes() ([]byte, error) {
	return json.Marshal(j)
}
