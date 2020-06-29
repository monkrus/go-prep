package main

import (
	"errors"
	"flag"
	"fmt"
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
	env = flag.String("env", "development", "a string")
	port = flag.Int("port", 8081, "an int")
	flag.Parse() //You forgot this command!
}

func helloLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello World</b>")

}

func main() {
	//here you initialized Gorilla's router
	router := mux.NewRouter()
	//then you used it for described below line
	router.HandleFunc("/hello", helloLink).Methods("GET")
	// create a postform
	router.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		if value, ok := query["name"]; ok {
			if err := WrongType(value[0]); err != nil {
				fmt.Fprintf(w, err.Error())
				//log.Fatal(err)
			} else {
				fmt.Fprintf(w, "Name: %s", value[0])
			}
		} else {
			fmt.Fprintf(w, "Oh, I don't know your name, stranger")
		}
	})
	http.ListenAndServe(":"+strconv.Itoa(*port), router) //You left string "8081" and didn't use port value!!!
}

//Solution 1

func WrongType(value string) error {
	//try co convert => if err is nil, it's an int => error
	if _, err := strconv.Atoi(value); err == nil {
		return errors.New("R u serious? No, R2D2 is OK, but " + value + "...")
	} else {
		return nil
	}
}

//Attempt 1

/* func wrongType(value string)error {
	if value ==  int { /*!!!! GO IS LANGUAGE WITHOUT FULL DUCK-TYPING, SUCH AS JS. YOU CAN'T COMPARE VALUE AND TYPE VIA "=="
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
