package controllers

import (
	"MyGream/database"
	"MyGream/helpers"
	"MyGream/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))
	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)

	} else {
		c.ShouldBind(&SocialMedia)

	}
	SocialMedia.UserId = userID
	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated,
		SocialMedia)
}
