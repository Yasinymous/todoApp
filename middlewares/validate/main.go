package validate

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
)

// func requestTime(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, "requestTime", time.Now().Format(time.RFC3339))
// 		r = r.WithContext(ctx)
// 		next.ServeHTTP(w, r)
// 	})
// }

func SignUpVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		userSignUp := &u.UserSignUp{}

		errbody := json.NewDecoder(req.Body).Decode(userSignUp)
		if errbody != nil {
			h.Respond(res, h.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
			return
		}

		if userSignUp.Password != userSignUp.RePassword {
			h.Respond(res, h.Message(false, "Şifreler Eslesmiyor!"))
			return
		}

		if len(userSignUp.Password) < 8 {
			h.Respond(res, h.Message(false, "Şifreniz en az 8 karakter olmalıdır!"))
			return
		}

		if !strings.Contains(userSignUp.Email, "@") {
			h.Respond(res, h.Message(false, "Email adresi hatalıdır!"))
			return
		}

		temp := &u.User{}

		errDb := db.GetDB().Table("users").Where("email = ?", userSignUp.Email).First(temp).Error
		if errDb != nil && errDb != gorm.ErrRecordNotFound {
			h.Respond(res, h.Message(false, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!"))
			return
		}

		if temp.Email != "" {
			h.Respond(res, h.Message(false, "Email adresi başka bir kullanıcı tarafından kullanılıyor."))
			return
		}
		user := &u.User{}

		user.Username = userSignUp.Username
		user.Email = userSignUp.Email
		user.Password = userSignUp.Password

		context.Set(req, "User", *user)
		next.ServeHTTP(res, req)
	})
}

// func (user *User) PasswordController() (map[string]interface{}, bool) {

// 	if len(user.Password) < 8 {
// 		return h.Message(false, "Şifreniz en az 8 karakter olmalıdır!"), false
// 	}

// 	return h.Message(false, "Her şey yolunda!"), true
// }

// func middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
// 		if len(authHeader) != 2 {
// 			fmt.Println("Malformed token")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Malformed Token"))
// 		} else {
// 			jwtToken := authHeader[1]
// 			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 				}
// 				return []byte(SECRETKEY), nil
// 			})

// 			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 				ctx := context.WithValue(r.Context(), "props", claims)
// 				// Access context values in handlers like this
// 				// props, _ := r.Context().Value("props").(jwt.MapClaims)
// 				next.ServeHTTP(w, r.WithContext(ctx))
// 			} else {
// 				fmt.Println(err)
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte("Unauthorized"))
// 			}
// 		}
// 	})
// }
