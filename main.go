package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

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
	{ //getting a 'too many colons' error
		hostName := "localhost:8081"
		portNum := "8081"
		seconds := 5
		timeOut := time.Duration(seconds) * time.Second

		conn, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", hostName, int64(seconds))
		fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
		fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())

	}
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
