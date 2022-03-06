package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	h "TaskManagement/helpers"
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

func (user *User) Create() map[string]interface{} {

	// if resp, ok := user.Validate(); !ok {
	// 	return resp
	// }
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return h.Message(false, "Bağlantı hatası oluştu. Kullanıcı yaratılamadı!")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.SessionToken = tokenString

	user.Password = ""

	response := h.Message(true, "Hesap başarıyla yaratıldı!")
	response["user"] = user
	return response
}

func GetUser(u int) map[string]interface{} {
	acc := &User{}
	GetDB().Table("users").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}

	acc.Password = ""
	response := h.Message(true, "Kullanici Bulundu")
	response["user"] = acc
	return response
}
