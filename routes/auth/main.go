package AuthRoutes

import (
	controllers "TaskManagement/controllers/auth"
	middlewares "TaskManagement/middlewares/validate"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/sign/in", controllers.SignIn).Methods("POST")
	r.Handle("/sign/up", middlewares.SignUpVerify(http.HandlerFunc(controllers.SignUp))).Methods("POST")
	r.HandleFunc("/sign/out", controllers.SignOut).Methods("POST")
}
