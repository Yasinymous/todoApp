package AuthRoutes

import (
	controllers "TaskManagement/controllers/auth"
	middlewares "TaskManagement/middlewares/validate"
	"net/http"

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
	r.HandleFunc("/sign/in", controllers.SignIn).Methods("POST")
	r.Handle("/sign/up", middlewares.SignUpVerify(http.HandlerFunc(controllers.SignUp))).Methods("POST")
	r.HandleFunc("/sign/out", controllers.SignOut).Methods("POST")
}
