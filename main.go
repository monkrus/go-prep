package main

import (
	"fmt"
	"net/http"
	"net/url"

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
		//CHOOSE ONE (BUT FIRST MORE AGILE)
		//FIRST VARIANT
		query, _ := url.ParseQuery(r.URL.RawQuery)
		if value, ok := query["name"]; ok {
			fmt.Fprintf(w, "Name: %s", value[0])
		} else {
			fmt.Fprintf(w, "Oh, i don't know your name, stranger")
		}
		//SECOND VARIANT
		query2 := r.URL.Query().Get("name")
		if query2 != "" {
			fmt.Fprintf(w, "Name: %s", query2)
		} else {
			fmt.Fprintf(w, "Oh, i don't know your name, stranger")
		}
	})
	//as second parametr you set "router" - it meens that ListenAndServe method accepted only "router" method
	http.ListenAndServe(":8081", router)

}
