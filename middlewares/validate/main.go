package validate

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	*u.User
}

func (user *User) EmailController() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return h.Message(false, "Email adresi hatalıdır!"), false
	}
	temp := &u.User{}

	err := db.GetDB().Table("accounts").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return h.Message(false, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!"), false
	}
	if temp.Email != "" {
		return h.Message(false, "Email adresi başka bir kullanıcı tarafından kullanılıyor."), false
	}

	return h.Message(false, "Her şey yolunda!"), true
}

func (user *User) PasswordController() (map[string]interface{}, bool) {

	if len(user.Password) < 8 {
		return h.Message(false, "Şifreniz en az 8 karakter olmalıdır!"), false
	}

	return h.Message(false, "Her şey yolunda!"), true
}
