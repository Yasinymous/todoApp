package UserController

import (
	h "TaskManagement/helpers"
	u "TaskManagement/models"
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
	resp := u.GetUser(i)
	h.Respond(res, resp)
}
