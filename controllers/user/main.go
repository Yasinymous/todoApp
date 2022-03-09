package UserController

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	u "TaskManagement/models/user"
	"fmt"
	"strconv"

	// "encoding/json"
	"net/http"
)

func GetAll(res http.ResponseWriter, req *http.Request) {
	var users []u.User

	result := db.GetDB().Table("users").Find(&users)

	response := h.Message(true, "Kullanicicilar Listelendi!")
	response["data"] = result.Value

	h.Respond(res, response)
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, ok := req.URL.Query()["userId"]

	if !ok || len(userId[0]) < 1 {
		h.Respond(res, h.Message(false, "Url Param 'key' is missing!"))
		return
	}

	i, err := strconv.Atoi(userId[0])
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	user := &u.User{}
	db.GetDB().Table("users").Where("id = ?", i).First(user)
	if user.Email == "" {
		h.Respond(res, h.Message(false, "Kullanici Bulunamadi!"))
		return
	}

	user.Password = ""
	response := h.Message(true, "Kullanici Bulundu")
	response["data"] = user

	h.Respond(res, response)
}
