package models

import "time"

type GormModel struct {
	ID       uint       `gorm:"primaryKey" json:"id"`
	CreateAt *time.Time `json:"create_at,omitempety"`
	UpdateAt *time.Time `json:"update_at,omitempety"`
}
