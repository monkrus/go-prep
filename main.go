package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

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

var (
	env  *string
	port *int
)

func init() {
	env = flag.String("env", "development", "a string")
	port = flag.Int("port", 8081, "an int")
}

func helloLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello World</b>")

}
/* Attmpt 2
func checkValue() ( value type, error) {
	values, err := values.type()
	if err!= nil {
    return nil, errors.Wrap(err, "this is a wrong type")
	}
*/

func main() {

	flag.Parse()
	fmt.Println("env:", *env)
	fmt.Println("port:", *port)

	//initialize Gorilla/Mux router
	router := mux.NewRouter()
	// utilize router below
	router.HandleFunc("/hello", helloLink).Methods("GET")
	// create a postform
	router.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		if value, ok := query["name"]; ok {
			fmt.Fprintf(w, "Name: %s", value[0])
		} else {
			fmt.Fprintf(w, "Oh, I don't know your name, stranger")
		}
                // Attempt 1
		if value = []int {
		fmt.Fprintf(w, "Int type is not allowed")}
		})
		
	  http.ListenAndServe(":8081", router)
         }

