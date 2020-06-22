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
	//here you initialized Gorilla's router
	router := mux.NewRouter()
	//then you used it for described below line
	router.HandleFunc("/hello", helloLink).Methods("GET")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./sergei.html")
	})
	router.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("username")
		fmt.Fprintf(w, "Name: %s", name)
	})
	//as second parametr you set "router" - it meens that ListenAndServe method accepted only "router" method
	http.ListenAndServe(":8081", router)

}
