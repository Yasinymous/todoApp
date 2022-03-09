package BoardRoutes

import (
	controllers "TaskManagement/controllers/board"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {

	r.StrictSlash(true)
	r.HandleFunc("/get/all", controllers.GetAll).Methods("GET")
	r.HandleFunc("/get", controllers.GetBoard).Methods("GET")
	r.HandleFunc("/add", controllers.AddBoard).Methods("POST")
	// r.HandleFunc("/set", controllers.SetBoard).Methods("PUT")
	// r.HandleFunc("/delete", controllers.DeleteBoard).Methods("DELETE")

}
