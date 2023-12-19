package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Simulated in-memory storage for URLs
var urlStorage = make(map[string]string)
var visitsCount = make(map[string]int)

func ShortenURL(c *gin.Context) {
	var requestBody struct {
		OriginalURL string `json:"original_url" binding:"required"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Generate a unique slug
	newSlug := generateSlug()
	urlStorage[newSlug] = requestBody.OriginalURL

	c.JSON(http.StatusCreated, gin.H{"shortened_url": newSlug})
}

func ModifySlug(c *gin.Context) {
	oldSlug := c.Param("slug")

	var requestBody struct {
		NewSlug string `json:"new_slug" binding:"required"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if the new slug already exists
	if _, exists := urlStorage[requestBody.NewSlug]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slug already exists"})
		return
	}

	// Update the slug
	urlStorage[requestBody.NewSlug] = urlStorage[oldSlug]
	delete(urlStorage, oldSlug)

	c.JSON(http.StatusOK, gin.H{"message": "Slug updated successfully"})
}

func RedirectToURL(c *gin.Context) {
	slug := c.Param("slug")
	originalURL, exists := urlStorage[slug]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Increment visit count
	visitsCount[slug]++
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func GetShortURLs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"short_urls": urlStorage})
}

// Generate unique slug
func generateSlug() string {
	rand.Seed(time.Now().UnixNano())

	timestamp := time.Now().UnixNano()

	randomBytes := make([]byte, 6)
	for i := range randomBytes {
		randomBytes[i] = byte(rand.Intn(26) + 65)
	}

	slug := strconv.FormatInt(timestamp, 10) + string(randomBytes)
	return slug
}
