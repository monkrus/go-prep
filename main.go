package main

import (
	"fmt"
	"log"
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
func helloSergei(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello Sergei</b>")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloLink).Methods("GET")
	router.HandleFunc("/hello/{id}", helloSergei).Methods("GET")
	http.ListenAndServe(":8081", router)

	addr := ":8081"
	log.Fatal(http.ListenAndServe(addr, nil))
}
