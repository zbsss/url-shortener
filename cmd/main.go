package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zbsss/url-shortener/shortener"
)

type shortenUrlRequest struct {
	OriginalUrl *string `json:"originalUrl"`
	ShortUrl    *string `json:"shortUrl"`
}

func main() {
	r := gin.Default()
	us := shortener.NewUrlShortener()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/urls", func(c *gin.Context) {
		var body shortenUrlRequest

		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatus(400)
		}

		var (
			shortUrl string
			err      error
		)
		if body.ShortUrl != nil {
			shortUrl, err = us.CreateCustomShortenedUrl(*body.OriginalUrl, *body.ShortUrl)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{
					"error": "this url is already taken",
				})
				return
			}
		} else {
			shortUrl = us.CreateRandomShortenedUrl(*body.OriginalUrl)
		}

		c.JSON(http.StatusCreated, gin.H{
			"shortUrl": shortUrl,
		})
	})

	r.GET("/urls/:shortUrl", func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")

		originalUrl, ok := us.GetOriginalUrl(shortUrl)

		if !ok {
			c.AbortWithStatus(404)
		}

		c.Redirect(http.StatusPermanentRedirect, originalUrl)
	})

	r.Run()
}
