package models

import "time"

type GormModel struct {
	ID       uint       `gorm:"primaryKey" json:"id" example:"1"`
	CreateAt *time.Time `json:"create_at,omitempety" example:"12 Feb 2023"`
	UpdateAt *time.Time `json:"update_at,omitempety" example:"12 Feb 2023"`
}
