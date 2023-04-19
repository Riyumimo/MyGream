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
// @Summary      Membuat Photo
// @Description  POST photo untuk mendapatkan user id ada di response dari login
// @Tags         CREATE Photo
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Success      200  {object}  models.Photo
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /photos/ [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	if contentType == appJson {
		c.ShouldBindJSON(&Photo)

	} else {
		c.ShouldBind(&Photo)

	}
	Photo.UserId = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)

}

// ShowAccount godoc
// @Summary      mengupdate Photo
// @Description  PUT photo untuk mendapatkan user id ada di response dari login
// @Tags         UPDATE PHOTO
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Success      200  {object}  models.Photo
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /photos/ [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	PhotoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	if contentType == appJson {
		c.ShouldBindJSON(&Photo)

	} else {
		c.ShouldBind(&Photo)

	}
	Photo.UserId = userID
	Photo.ID = uint(PhotoId)

	err := db.Model(&Photo).Where("id = ?", PhotoId).Updates(models.Photo{
		Title:     Photo.Title,
		Photo_Url: Photo.Photo_Url,
		Caption:   Photo.Caption,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

// ShowAccount godoc
// @Summary      mendapatkan semua Photo
// @Description  GETALL photo
// @Tags         GEL ALL PHOTO
// @Accept       json
// @Produce      json
// @Success      200  {array}  []models.Photo
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /photos/ [get]
func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	var photos []models.Photo
	var comments []models.Comment
	if err := db.Order("id asc").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	err := db.Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	for _, c := range comments {
		for i, v := range photos {
			if v.ID == c.PhotoId {
				photos[i].Comment = append(photos[i].Comment, c)
				break
			}
		}
	}

	c.JSON(http.StatusOK, photos)

}

// ShowAccount godoc
// @Summary      mendapatkan Photo bersarkan dari id nya
// @Description  GET photo menggunakan id sebagai param
// @Tags         GET PHOTO by ID
// @Accept       json
// @Produce      json
// @Param  		id header string true "id {id}"
// @Success      200  {object}  models.Photo
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /photos/{id} [get]
func GetPhotoByID(c *gin.Context) {
	db := database.GetDB()

	photoID := c.Param("id")

	var photo models.Photo
	err := db.Preload("Comment").Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          photo.ID,
		"title":       photo.Title,
		"caption":     photo.Caption,
		"photo_url":   photo.Photo_Url,
		"user_id":     photo.UserId,
		"create_date": photo.CreateAt,
		"update_date": photo.UpdateAt,
		"comments":    photo.Comment,
	})
}

// ShowAccount godoc
// @Summary      mendelete Photo berdsarkan id nya
// @Description  DELETE photo berdasarkan id dan untuk mendapatkan user id ada di response dari login
// @Tags         DELETE PHOTO by ID
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param id header string true "id {id}"
// @Success      200  {object}  models.Photo
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /photos/{id} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoID := c.Param("id")
	userID := uint(userData["id"].(float64))

	//mencari foto yang akan di hapus

	var photo models.Photo
	var comment models.Comment

	if err := db.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}
	if err := db.First(&comment, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	//memastikan userid sama dengan photo userid
	if photo.UserId != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this photo"})
		return
	}

	//menghapus photo dari databse

	if err := db.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete photo"})
		return
	}
	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete photo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "photo deleted successfully"})

}
