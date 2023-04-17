package controllers

import (
	"MyGream/database"
	"MyGream/helpers"
	"MyGream/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photoID := c.Param("id")

	comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	if contentType == appJson {
		c.ShouldBindJSON(&comment)

	} else {
		c.ShouldBind(&comment)

	}
	id, _ := strconv.Atoi(photoID)
	comment.PhotoId = uint(id)
	comment.UserId = userID
	err := db.Debug().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated,
		comment)
}
