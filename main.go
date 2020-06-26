package main

import (
	
	"flag"
	"fmt"
	"log"
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
	})    
	  http.ListenAndServe(":8081", router)
	}
	//Attempt 1
	
	/* func wrongType(value string)error {
		if value ==  int {
			return fmt.Errorf("int type is not allowed")
		}
	//Attempt 2

	var  (
		wrongType =  fmt.Errorf("int type is not allowed")
	)
		
	func wrongType(value string) error {
		if value ==  int {
			return fmt.Errorf("int type is not allowed")
		}
	//	Attempt 3

	var wrongType =errors.New("int type is not allowed")
	func checkType(value string) (string, error){
		if value == int {
			return "", wrongType
		}
	}


    */

		}
	}