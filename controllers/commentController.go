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

// ErrorResponse response struct
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ShowAccount godoc
// @Summary      Membuat comment
// @Description  POST comment by photo id dan untuk mendapatkan user id menggunkan token bearer yang di dapat dari login
// @Tags         CREATE COMMENT
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param        PhotoId   path      int  true  "Photo Id"
// @Success      200  {object}  models.Comment
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /comment/photo_id{id} [post]
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

// ShowAccount godoc
// @Summary      Get all data comment
// @Description  Get semua data comment
// @Tags         GET ALL COMMENT
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Success      200  {object}  []models.Comment
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /comment/ [get]
func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	var comments []models.Comment
	if err := db.Order("id asc").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// ShowAccount godoc
// @Summary      Get comment by id
// @Description  GET comment by  id
// @Tags         GET COMMENT by id
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param        Id   path      int  true  "Id"
// @Success      200  {object}  models.Comment
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /comment/{id} [get]
func GetCommentByID(c *gin.Context) {
	db := database.GetDB()

	Id := c.Param("id")
	var comments models.Comment
	err := db.Where("id = ?", Id).First(&comments).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          comments.ID,
		"user_id":     comments.UserId,
		"create_date": comments.CreateAt,
		"update_date": comments.UpdateAt,
		"message":     comments.Message,
		"PhotoId":     comments.PhotoId,
	})
}

// ShowAccount godoc
// @Summary      Mengupdata data comment yang telah ada
// @Description  PUT comment by photo id dan untuk mendapatkan user id menggunkan token bearer yang di dapat dari login
// @Tags         PUT COMMENT
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param        Id   path      int  true  "Id"
// @Success      200  {object}  models.Comment
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /comment/{id} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	CommentId, _ := strconv.Atoi(c.Param("id"))
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	if contentType == appJson {
		c.ShouldBindJSON(&Comment)

	} else {
		c.ShouldBind(&Comment)

	}
	Comment.UserId = userID
	Comment.ID = uint(CommentId)
	err := db.Model(&Comment).Where("id = ?", CommentId).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

// ShowAccount godoc
// @Summary      Delete Comment yang telah di buat
// @Description  Delete comment by  id dan untuk mendapatkan user id menggunkan token bearer yang di dapat dari login
// @Tags         DELETE COMMENT
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer {token}"
// @Param        Id   path      int  true  "Id"
// @Success      200  {object}  models.Comment
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	commetnID := c.Param("id")
	userID := uint(userData["id"].(float64))
	var comment models.Comment

	//cari commnet yang akan di hapus
	if err := db.First(&comment, commetnID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	//memastikan userid sama dengan commnent userid
	if comment.UserId != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this photo"})
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment deleted successfully"})

}
