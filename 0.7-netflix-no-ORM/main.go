package main

import (
	"fmt"
	"net/http"

	"main.go/routers"
)

func main() {

	fmt.Println("MongoDB API")

	router := routers.Router()
	fmt.Println("\nServer is getting started...")

	http.ListenAndServe(":4000", router)
	fmt.Println("\nListening at port 4000 ...")
}
