// main.go
package main

import (
	middlewares "TaskManagement/middlewares/jwt"
	AuthRoutes "TaskManagement/routes/auth"
	BoardRoutes "TaskManagement/routes/board"
	UserRoutes "TaskManagement/routes/user"

	// BoardRoutes "TaskManagement/routes/board"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// srv := &http.Server{
	// 	Handler: AuthRoutes.AuthRoutes(),
	// 	Addr:    "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	r := mux.NewRouter()
	r.Use(middlewares.JwtAuthentication)
	authRoutes := r.PathPrefix("/auth").Subrouter()
	userRoutes := r.PathPrefix("/user").Subrouter()
	boardRoutes := r.PathPrefix("/board").Subrouter()
	AuthRoutes.RegisterHandlers(authRoutes)
	UserRoutes.RegisterHandlers(userRoutes)
	BoardRoutes.RegisterHandlers(boardRoutes)

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Serving error.")
	}
}
