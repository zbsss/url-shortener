package shortener

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	UrlShortener interface {
		CreateRandomShortenedUrl(longUrl string) string
		CreateCustomShortenedUrl(longUrl, shortenedUrl string) (string, error)
		GetOriginalUrl(shortUrl string) (string, bool)
	}

	urlShortener struct {
		urls map[string]string
	}
)

func NewUrlShortener() UrlShortener {
	return &urlShortener{
		urls: make(map[string]string),
	}
}

func (us *urlShortener) CreateRandomShortenedUrl(longUrl string) string {
	for {
		shortened := uuid.New().String()

		_, alreadyExists := us.urls[shortened]

		if !alreadyExists {
			us.urls[shortened] = longUrl
			return shortened
		}
	}
}

func (us *urlShortener) CreateCustomShortenedUrl(longUrl, shortUrl string) (string, error) {
	_, alreadyExists := us.urls[shortUrl]

	if alreadyExists {
		return "", fmt.Errorf("this short url is already in use")
	}

	us.urls[shortUrl] = longUrl

	return shortUrl, nil
}

func (us *urlShortener) GetOriginalUrl(shortUrl string) (string, bool) {
	value, ok := us.urls[shortUrl]
	return value, ok
}
