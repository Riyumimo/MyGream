package models

import (
	"MyGream/helpers"
	"errors"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	UserName    string        `gorm:"not null;uniqueIndex" json:"full_name" form:"full_name" valid:"required-Your full name is required" example:"ExampleJunior"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required-Your email is required,email-Invalid email format" example:"example@gmail.com"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required-Your password is required,min=6,Password has to have minimum length of 6 characters" example:"1234567"`
	Age         string        `gorm:"not null" json:"age" form:"age" valid:"required-Your age is required,min=9,Password has to have minimum length of 8 year" example:"12"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedias"`
	// Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if err != nil {
		err = errCreate
		return
	}

	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	u.Password = helpers.HassPass(u.Password)
	if helpers.IsEmailValid(u.Email) == false {
		return errors.New("email is not valid")
	}
	ageStr := u.Age
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return
	}

	if age < 8 {
		return errors.New("Age must be at least 9 years old ")
	}
	return
}
