package routes

import (
	"Golang/controllers"
	"Golang/utils"
	"net/http"
)

func RegisterRoutes(r *http.ServeMux) {

	r.HandleFunc("POST /register",controllers.RegisterUser)
	r.HandleFunc("POST /login", controllers.LoginUser)


   // Secure routes with JWT middleware
   r.Handle("GET /todos", utils.JWTMiddleware(http.HandlerFunc(controllers.GetTodos)))
   r.Handle("POST /todos", utils.JWTMiddleware(http.HandlerFunc(controllers.CreateTodos)))
   r.Handle("DELETE /todos/{ID}", utils.JWTMiddleware(http.HandlerFunc(controllers.DeleteTodos)))
   r.Handle("PUT /todos/{ID}", utils.JWTMiddleware(http.HandlerFunc(controllers.UpdateTodos)))
}