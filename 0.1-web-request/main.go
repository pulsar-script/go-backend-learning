package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/10"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// Response type
	fmt.Printf("The type of Response is => %T", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// print response content

	content := string(databytes)
	fmt.Println(content)

}
