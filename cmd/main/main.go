package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// absolute path routes importation
	"github.com/mkyd-kill/Bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// routers handler
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}