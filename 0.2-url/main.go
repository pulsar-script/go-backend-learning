package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://hero.dev:3000/learn?user=monkeymagic&&rank=hero"

func main() {

	fmt.Println("URL tutorial")
	fmt.Println(myURL) // normal printing

	//* parsing url | cutting into small parts with means
	result, err := url.Parse(myURL)

	if err != nil {
		panic(err)
	}

	fmt.Println(" result => ", result)
	fmt.Println(" result Schema=> ", result.Scheme)
	fmt.Println(" result Host => ", result.Host)
	fmt.Println(" result Path => ", result.Path)
	fmt.Println(" result Port()  => ", result.Port()) // Port() is method
	fmt.Println(" result => ", result.RawQuery)       // params

	// better way to seperate / extarct query parameters
	queryParamsResult := result.Query()
	fmt.Printf("Tyepe of Query() result => %T", queryParamsResult)
	// output => url.Values basically key values

	fmt.Println("")
	fmt.Println(queryParamsResult["user"])
	fmt.Println(queryParamsResult["rank"])

	// loop

	fmt.Println("")
	for queryName, value := range queryParamsResult {
		fmt.Printf(" %v => %v \n", queryName, value)
	}

	//* Constructing URL
	partsOfUrl := &url.URL{ // we have to pass original value address, not copy
		Scheme:  "https",
		Host:    "kai.uk",
		Path:    "/anime",
		RawPath: "name=overlord",
	}

	anotherMyUrl := partsOfUrl.String() // similar as string(partOfUrl)

	fmt.Println("constructed URL => ", anotherMyUrl)
}
