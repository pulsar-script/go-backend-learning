package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request tutorial of Go Backend Dev")

	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

// func to make get request
func PerformGetRequest() {
	const myUrl = "http://localhost:3000/hello"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // callers responsibility to close the connection after user
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	// content in body
	content, _ := io.ReadAll(response.Body)

	//* 1st way - easy
	fmt.Println("Content in bytes : ", content)
	fmt.Println("Content : ", string(content))

	// 2nd way (recommended way)
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("byte count : ", byteCount)
	fmt.Println(" responseString (builder) : ", responseString) //* This always hold original value into bytes , and due to its build-in methods we can convert into many format
	fmt.Println("content into string format : ", responseString.String())
}

// func to make POST Json request
func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/send/message"

	// fake json payload
	bodyContent := strings.NewReader(`

		{
			"username" : "heroHamada",
			"age" : 21,
			"role" : "manipulator"
		}

		`)

	// call
	response, err := http.Post(myUrl, "application/json", bodyContent)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // callers responsibility to close the connection after user
	content, _ := io.ReadAll(response.Body)
	fmt.Println("content from POST request : ", string(content))

}

// func to make POST Form request
func PerformPostFormRequest() {

	const myUrl = "http://localhost:3000/send/form"

	//form data - this is special format which HTML forms use to send data

	data := url.Values{}

	data.Add("firstname", "hero")
	data.Add("lastname", "hamada")
	data.Add("username", "king")

	// call
	response, err := http.PostForm(myUrl, data) //* in form-data case in backend we dont get data in json format , thats why at the time of sending to frontend we have to convert it into json
	//*  unlike, sending above non - from data / json data we pass to backend

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // callers responsibility to close the connection after user

	content, _ := io.ReadAll(response.Body)

	fmt.Println("content from Form Post request : ", string(content))

}
