package product

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Route struct {
	Path        string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

func GetRoutes() []Route {

	getall := Route{Path: "/product", HandlerFunc: getAll}
	getp := Route{Path: "/product/{id}", HandlerFunc: getp}

	routes := []Route{getall, getp}
	return routes

}

func getAll(w http.ResponseWriter, r *http.Request) {

	data, err := json.MarshalIndent(Products, "\n", "\t")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("ContentType", "application/json")
	w.Header().Add("ContentLenght", fmt.Sprint(len(data)))
	w.WriteHeader(http.StatusOK)
	n, err := w.Write(data)

	if err != nil {
		fmt.Println("Received error")
	}
	fmt.Println("wrote", n, "bytes")
}

func getp(w http.ResponseWriter, r *http.Request) {
}
