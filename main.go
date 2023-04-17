package main

import (
	"go-website/config"
	"go-website/controllers/categorycontroller"
	"go-website/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Starting server...")
	http.ListenAndServe(":8081", nil)

}
