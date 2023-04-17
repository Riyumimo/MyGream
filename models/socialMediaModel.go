package models

import (
	// "MyGram/vendor/github.com/asaskevich/govalidat

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required-Your Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required-Your social_media_url is required"`
	UserId         uint
	User           *User
	// Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if err != nil {
		err = errCreate
		return
	}
	err = nil
	return

}

func (p *SocialMedia) AfterCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if err != nil {
		err = errCreate
		return
	}
	err = nil
	return

}
