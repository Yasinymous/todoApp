package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	// Id           int       `gorm:"not null" json:"Id"`
	Username     string `gorm:"not null" json:"Username"`
	Email        string `gorm:"not null" json:"Email"`
	Password     string `gorm:"not null" json:"Password"`
	SessionToken string `json:"SessionToken"`
}

type Token struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}
