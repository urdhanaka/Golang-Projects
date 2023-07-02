package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type URL struct {
	ID           int    `json:"id"`
	Url          string `json:"url"`
	ShortenedUrl string `json:"short"`
}

var urlList []URL
var lastID int

func main() {
	r := gin.Default()

	// GET Method
	r.GET("/", helloWorld)
	r.GET("/:url", getOriginalUrl)

	// POST Method
	r.POST("/", addNewUrl)

	r.Run(":8080")
}

func helloWorld(c *gin.Context) {
	c.JSON(200, urlList)
}

func getOriginalUrl(c *gin.Context) {
	shortUrlString := c.Param("url")

	for _, urlStruct := range urlList {
		shortUrl := urlStruct.ShortenedUrl

		if shortUrl == shortUrlString {
			c.JSON(http.StatusOK, urlStruct)
			return
		}
	}

	c.JSON(http.StatusNotFound, urlList)
}

func addNewUrl(c *gin.Context) {
	var urlStruct URL

	if err := c.ShouldBindJSON(&urlStruct); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	lastID++
	urlStruct.ID = lastID
	urlList = append(urlList, urlStruct)

	c.JSON(201, urlList)
}
