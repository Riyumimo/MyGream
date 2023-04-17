package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required-Your Title is required"`
	Caption   string `gorm:"not null" json:"caption" form:"caption" valid:"required-Your caption is required"`
	Photo_Url string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required-Your photo_url is required"`
	UserId    uint
	User      *User
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	// Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if err != nil {
		err = errCreate
		return
	}
	err = nil
	return

}

func (p *Photo) AfterCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if err != nil {
		err = errCreate
		return
	}
	err = nil
	return

}
