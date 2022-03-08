package AuthController

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}
	temp := &u.User{}
	errbody := json.NewDecoder(req.Body).Decode(temp)
	if errbody != nil {
		h.Respond(res, h.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
		return
	}

	err := db.GetDB().Table("users").Where("username = ?", temp.Username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			h.Respond(res, h.Message(false, "Kullanici adi bulunamadı!"))
			return
		}
		h.Respond(res, h.Message(false, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(temp.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { // Parola eşleşmedi
		h.Respond(res, h.Message(false, "Parola hatalı! Lütfen tekrar deneyiniz!"))
		return
	}

	// Giriş başarılı
	user.Password = ""

	// JWT yaratılır
	tk := &u.Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	user.SessionToken = tokenString // JWT yanıta eklenir
	db.GetDB().Model(&u.User{}).Where("ID = ?", user.ID).Update("session_token", user.SessionToken)

	response := h.Message(true, "Giriş başarılı!")
	response["data"] = user
	h.Respond(res, response)
}
func SignUp(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}
	User := context.Get(req, "User")

	if u, ok := User.(u.User); ok {
		user = &u
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db.GetDB().Create(user)

	if user.ID <= 0 {
		h.Respond(res, h.Message(false, "Bağlantı hatası oluştu. Kullanıcı yaratılamadı!"))
		return
	}

	tk := &u.Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.SessionToken = tokenString

	user.Password = ""

	response := h.Message(true, "Hesap başarıyla yaratıldı!")
	response["user"] = user

	h.Respond(res, response)
}
func SignOut(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}
	User := context.Get(req, "User")
	if u, ok := User.(u.User); ok {
		user = &u
	}

	db.GetDB().Model(&u.User{}).Where("ID = ?", user.ID).Update("session_token", nil)
	response := h.Message(true, "Cikis başarılı!")
	response["data"] = user
	h.Respond(res, response)
}
