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

// ShowAccount godoc
// @Summary      membuat SocialMedia
// @Description  POST SocialMedia berdasarkan user id dan untuk mendapatkan user id ada di response dari login
// @Tags         CREATE SOCIAL MEDIA
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Success      200  {object}  models.SocialMedia
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /socialmedia/ [post]
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

// ShowAccount godoc
// @Summary      mendapatkan semua SocialMedia
// @Description  GETALL semua  SocialMedia
// @Tags         GETALL SOCIAL MEDIA
// @Accept       json
// @Produce      json
// @Success      200  {array}  []models.SocialMedia
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /socialmedia/ [get]
func GetAllSocial(c *gin.Context) {
	db := database.GetDB()
	var social []models.SocialMedia

	if err := db.Order("id asc").Find(&social).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, social)

}

// ShowAccount godoc
// @Summary      mendapatkan SocialMedia bersarkan id ya
// @Description  GET SocialMedia berdasarkan id
// @Tags         GET SOCIAL MEDIA by ID
// @Accept       json
// @Produce      json
// @Param 		id header string true "id {id}"
// @Success      200  {object}  models.SocialMedia
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /socialmedia/{id} [get]
func GetSocialById(c *gin.Context) {
	db := database.GetDB()
	SocialId := c.Param("id")
	var social models.SocialMedia
	err := db.Where("id = ?", SocialId).First(&social).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Social Media Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, social)
}

// ShowAccount godoc
// @Summary      mengupdate social media SocialMedia berdasarkan id nya
// @Description  PUT SocialMedia berdasarkan id dan user id untuk mendapatkan user id ada di response dari login
// @Tags         UPDATE SOCIAL MEDIA
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param 		id header string true "id {id}"
// @Success      200  {object}  models.SocialMedia
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /socialmedia/{id} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socialID, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))
	Social := models.SocialMedia{}

	if contentType == appJson {
		c.ShouldBindJSON(&Social)

	} else {

		c.ShouldBind(&Social)

	}

	Social.ID = uint(socialID)
	Social.UserId = userID
	err := db.Model(&Social).Where("id = ?", socialID).Updates(models.SocialMedia{
		Name: Social.Name, SocialMediaUrl: Social.SocialMediaUrl,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Social)

}

// ShowAccount godoc
// @Summary      mendelete SocialMedia berdasarkan id
// @Description  DELETE SocialMedia berdasarkan id dan user id  untuk mendapatkan user id ada di response dari login
// @Tags         DELETE SOCIAL MEDIA
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param 	id path string true "id {token}"
// @Success      200  {object}  models.SocialMedia
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /socialmedia/{id} [delete]
func DeleteSocial(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	socialID := c.Param("id")
	userID := uint(userData["id"].(float64))
	var social models.SocialMedia

	//cari commnet yang akan di hapus
	if err := db.First(&social, socialID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "socialmedia not found"})
		return
	}

	//memastikan userid sama dengan commnent userid
	if social.UserId != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this photo"})
		return
	}

	if err := db.Delete(&social).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "socialmedia deleted successfully"})

}
