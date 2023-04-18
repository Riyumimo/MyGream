package models

import (
	// "MyGram/vendor/github.com/asaskevich/govalidat

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserId  uint
	PhotoId uint
	Message string `gorm:"not null" json:"message" form:"message" valid:"required-Your message is required"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if err != nil {
		err = errCreate
		return
	}

	err = nil
	return

}

func (p *Comment) AfterCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if err != nil {
		err = errCreate
		return
	}
	err = nil
	return

}
