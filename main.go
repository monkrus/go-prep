package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
//CREATE BASIC WEB SERVER
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<b>Hello World</b>")
	})
	addr := ":8081"
	log.Fatal(http.ListenAndServe(addr, nil))
}
*/

func helloLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello World</b>")

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloLink).Methods("GET")
	http.ListenAndServe(":8081", router)
	http.HandleFunc("/sergei", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, "Name: %s", name)
	})

}
