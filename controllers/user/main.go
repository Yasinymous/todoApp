package UserController

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"fmt"
	"log"
	"strconv"

	// "encoding/json"
	"net/http"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	// user := &u.User{}
	// err := json.NewDecoder(req.Body).Decode(user) // İstek gövdesi decode edilir, hatalı ise hata döndürülür
	// if err != nil {
	// 	h.Respond(res, h.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
	// 	return
	// }

	userId, ok := req.URL.Query()["userId"]

	if !ok || len(userId[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	i, err := strconv.Atoi(userId[0])
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	user := &u.User{}
	db.GetDB().Table("users").Where("id = ?", i).First(user)
	// if acc.Email == "" {
	// 	return nil
	// }

	user.Password = ""
	response := h.Message(true, "Kullanici Bulundu")
	response["user"] = user

	h.Respond(res, response)
}
