package repository

type Repository interface {
	CreateLink(shortLink, url string) error
	GetLink(shortLink string) (string, bool, error)
	LinkExists(shortLink string) (bool, error)
}
