package routes

// application routes here

import (
	"github.com/gorilla/mux"
	// absolute path import
	"github.com/mkyd-kill/Bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// router handler
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}