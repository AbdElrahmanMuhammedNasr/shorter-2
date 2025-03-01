package routes

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorter/db"
	"shorter/dto"
	"shorter/model"
)

func getUrls(context *gin.Context) {
	var urls []model.Urls
	result := db.DB.Find(&urls)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URLs"})
		return
	}

	context.JSON(http.StatusOK, urls)
}

func getUrl(context *gin.Context) {
	shortCode := context.Param("short_code")
	var url model.Urls
	result := db.DB.Where("short_code = ?", shortCode).First(&url)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	context.Redirect(http.StatusFound, url.OriginalURL)
}

func addUrl(context *gin.Context) {
	urlDTO := dto.UrlDTO{}
	if err := context.ShouldBindJSON(&urlDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//go func() {
	shortCode := generateShortCode()
	newURL := model.Urls{ShortCode: shortCode, OriginalURL: urlDTO.OriginalURL, UserID: urlDTO.UserID}
	db.DB.Create(&newURL)
	context.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortCode})
	//}()
}

func generateShortCode() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}
