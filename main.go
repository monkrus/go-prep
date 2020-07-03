package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

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
	//env = flag.String("env", "development", "a string")
	//port = flag.Int("port", 8081, "an int")
	//flag.Parse() 

}

func helloLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello World</b>")

}

func main() {
	//initialize Gorilla/Mux router
	router := mux.NewRouter()
	// utilize router below
	router.HandleFunc("/hello", helloLink).Methods("GET")
	// create a postform, define a query
	router.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		// define a value,name
		if value, ok := query["name"]; ok {
			if WrongType(value[0]) != nil {
				fmt.Fprintf(w, WrongType(value[0]).Error())
				//log.Fatal(err)
			} else {
				fmt.Fprintf(w, "Name: %s", value[0])
			}
		} else {
			fmt.Fprintf(w, "Oh, I don't know your name, stranger")
		}
	})
	// values from init function
	http.ListenAndServe(":"+strconv.Itoa(*port), router)

}

func WrongType(value string) error {
	//if err is nil =>(no error), then it's an int =>(error)
	if _, err := strconv.Atoi(value); err == nil {
		return errors.New("R u serious? No, R2D2 is OK, but " + value + "...")
	} else {
		return nil
	}
}
