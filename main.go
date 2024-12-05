package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kritimauludin/go-restapi-mux-gorm-mysql/controllers/productcontroller"
)

func main()  {
	
	routes := mux.NewRouter()

	routes.HandleFunc("/api/v1/products", productcontroller.Index).Methods("GET")
	routes.HandleFunc("/api/v1/product/{id}", productcontroller.Show).Methods("GET")
	routes.HandleFunc("/api/v1/product", productcontroller.Create).Methods("POST")
	routes.HandleFunc("/api/v1/product/{id}", productcontroller.Update).Methods("PUT")
	routes.HandleFunc("/api/v1/product", productcontroller.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", routes)
}