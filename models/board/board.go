package models

import (
	"github.com/jinzhu/gorm"
)

type Board struct {
	gorm.Model
	Title       string `gorm:"not null" json:"Title"`
	Description string `json:"Description"`
	UserId      int    `gorm:"not null"`
}
