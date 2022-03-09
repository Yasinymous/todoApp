package AuthRoutes

import (
	controllers "TaskManagement/controllers/user"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/get/all", controllers.GetAll).Methods("GET")
	r.HandleFunc("/get", controllers.GetUser).Methods("GET")
}
