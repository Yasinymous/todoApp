package AuthController

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the SignIn!")
	fmt.Println("Endpoint Hit: SignIn")
}
func SignUp(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}

	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		h.Respond(res, h.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
		return
	}

	// if resp, ok := user.Validate(); !ok {
	// 	return resp
	// }
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db.GetDB().Create(user)

	// if user.ID <= 0 {
	// 	return h.Message(false, "Bağlantı hatası oluştu. Kullanıcı yaratılamadı!")
	// }

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
