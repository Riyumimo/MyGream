package controllers

import (
	"MyGream/database"
	"MyGream/helpers"
	"MyGream/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

// ini adalah example token
type Token struct {
	Token string `json:"token" example:"U3dhZ2dlciByb2Nrc......."`
}

// ShowAccount godoc
// @Summary      Membuat Account
// @Description  POST account
// @Tags         CREATE ACCOUNT
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /register/ [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.UserName,
		"age":       User.Age,
		"create_at": time.Now().Format(time.UnixDate),
		"update_at": time.Now().Format(time.UnixDate),
	})

}

// ShowAccount godoc
// @Summary      Login Account
// @Description  POST login account
// @Tags         LOGIN ACCOUNT
// @Accept       json
// @Produce      json
// @Success      200  {object}  Token
// @Failure      400   {object} ErrorResponse
// @Failure 	 500 {object} ErrorResponse
// @Router       /login/ [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
