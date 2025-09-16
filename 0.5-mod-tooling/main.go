package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // <-- THIS IS THE MISSING LINE
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to golang series on YT </h1>"))
}
