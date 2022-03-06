package AuthController

import (
	h "TaskManagement/helpers"
	u "TaskManagement/models"
	"encoding/json"
	"fmt"
	"net/http"
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

	resp := user.Create()
	h.Respond(res, resp)
}
func SignOut(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the SignOut!")
	fmt.Println("Endpoint Hit: SignOut")
}
