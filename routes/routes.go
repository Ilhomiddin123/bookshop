package routes

import (
	"bookShop/pkg/controllers"
	"fmt"
	"net/http"
)

func RunRoutes() {

	http.HandleFunc("/user2s", controllers.GetAllUsers)
	http.HandleFunc("/users/getByID/", controllers.GetUserByID)
	http.HandleFunc("/users/create", controllers.CreateUser)
	http.HandleFunc("/users/editByID/", controllers.EditUserByID)
	http.HandleFunc("/users/deleteByID/", controllers.DeleteUserByID)

	http.HandleFunc("/books", controllers.GetAllBooks)
	http.HandleFunc("/books/getByID/", controllers.GetBookByID)
	http.HandleFunc("/books/create", controllers.CreateBook)
	http.HandleFunc("/books/editByID/", controllers.EditBookByID)
	http.HandleFunc("/books/deleteByID/", controllers.DeleteBookByID)

	fmt.Println("Server started")
	http.ListenAndServe(":1111", nil)
}
