package Jwt

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		response := make(map[string]interface{})
		notAuth := []string{"/auth/sign/in", "/auth/sign/up"} // Doğrulama istemeyen endpointler
		requestPath := req.URL.Path                           // mevcut istek yolu

		// Gelen isteğin doğrulama isteyip istemediği kontrol edilir
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(res, req)
				return
			}
		}

		tokenHeader := req.Header.Get("Authorization") // Header'dan token alınır

		if tokenHeader == "" { // Token yoksa "403 Unauthorized" hatası dönülür
			response = h.Message(false, "Token gönderilmelidir!")
			res.WriteHeader(http.StatusForbidden)
			res.Header().Add("Content-Type", "application/json")
			h.Respond(res, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") // Token'ın "Bearer {token} / Token {token}" formatında gelip gelmediği kontrol edilir
		if len(splitted) != 2 {
			response = h.Message(false, "Hatalı ya da geçersiz token!")
			res.WriteHeader(http.StatusForbidden)
			res.Header().Add("Content-Type", "application/json")
			h.Respond(res, response)
			return
		}

		tokenPart := splitted[1] // Token'ın doğrulama yapmamıza yarayan kısmı alınır

		tempToken := &u.User{}
		errToken := db.GetDB().Table("users").Where("session_token = ?", tokenPart).First(tempToken).Error
		if errToken != nil && errToken != gorm.ErrRecordNotFound {
			h.Respond(res, h.Message(false, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!"))
			return
		}
		if tempToken.SessionToken == "" {
			h.Respond(res, h.Message(false, "Boyle bir token kullanicida yok!"))
			return
		}

		tk := &u.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { // Token hatalı ise 403 hatası dönülür
			response = h.Message(false, "Token hatalı!")
			res.WriteHeader(http.StatusForbidden)
			res.Header().Add("Content-Type", "application/json")
			h.Respond(res, response)
			return
		}

		if !token.Valid { // Token geçersiz ise 403 hatası dönülür
			response = h.Message(false, "Token geçersiz!")
			res.WriteHeader(http.StatusForbidden)
			res.Header().Add("Content-Type", "application/json")
			h.Respond(res, response)
			return
		}

		// Doğrula başarılı ise işleme devam edilir

		user := &u.User{}

		user.ID = tk.UserId

		context.Set(req, "User", *user)
		next.ServeHTTP(res, req)
	})
}
