package main

import (
	"fmt"
	"net/http"
	"os"
	"quickcommerce/product"

	"github.com/gorilla/mux"
)

const (
	port = ":3040"
)

var (
	routes = []string{"/", "/product", "customer", "transaction"}
)

func main() {
	fmt.Println("Quick commerce...")

	//handler := http.NewServeMux()
	handler := mux.NewRouter()

	routes := product.GetRoutes()

	// prodcuts
	for _, r := range routes {
		handler.HandleFunc(r.Path, r.HandlerFunc).Methods("GET")
	}

	err := http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

}
