package AuthController

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the SignIn!")
	fmt.Println("Endpoint Hit: SignIn")
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
	fmt.Fprintf(res, "Welcome to the SignOut!")
	fmt.Println("Endpoint Hit: SignOut")
}
