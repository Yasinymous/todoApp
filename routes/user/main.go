package AuthRoutes

import (
	controllers "TaskManagement/controllers/user"

	"github.com/gorilla/mux"
)

// func AuthRoutes() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/sign/in", AuthController.SignIn).Methods("POST")
// 	r.HandleFunc("/sign/up", AuthController.SignUp).Methods("POST")
// 	r.HandleFunc("/sign/out", AuthController.SignOut).Methods("POST")
// 	return r
// }

func RegisterHandlers(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/get/", controllers.GetUser).Methods("GET")
	r.HandleFunc("/get/all", controllers.GetAll).Methods("GET")
}
