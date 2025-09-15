package main

import (
	"encoding/json"
	"fmt"
)

// struct for user
type user struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // this "-" exclude it from marshelling
	Age      int    `json:"age"`
	// Tags     []string `json:"tags, omitempty"` //* this "omitempty is exclude field when value is nil"
	Tags []string `json:"tags,omitempty"` // this "omitempty is exclude field when value is nil" //! carefully write synatx, dont give space => tags,omitempty"`
}

func main() {

	fmt.Println("Welcome to Mashelling Tutorial of Go")
	EncodeJson()
	DecodJson()
}

// func to Encode Data into Json , To send API suppose

func EncodeJson() {

	// Dummy data
	allUsers := []user{
		{"King76", "king12@gmail.com", "asbc12", 34, []string{"king", "Ruler"}},
		{"hamada6", "hamada12@gmail.com", "asbc12", 34, nil}, // nil value
		{"chocolate45", "chocolate12@gmail.com", "asbc12", 34, []string{"sweet", "Ruler"}},
	}

	//* package (pack) / convert this data into json data

	// with Marshal
	finalJsonNormal, err := json.Marshal(allUsers)

	if err != nil {
		panic(err)
	}

	// with MarshalIntend
	finalJsonIntent, err := json.MarshalIndent(allUsers, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n", finalJsonNormal)
	fmt.Printf("\n%s\n", finalJsonIntent)

}

// Decode Json , for Data come from API suppose and you want to consume as JSON

func DecodJson() {

	// all data come from API is in to bytes format
	jsonDataFromAPI := []byte(`  
        {
                "username": "King76",
                "email": "king12@gmail.com",
                "age": 34,
                "tags": ["king","Ruler"]
        }
	`)

	// before unmarshalling we need to check, is it a valid json ?
	checkValid := json.Valid(jsonDataFromAPI)

	//* 1ST CASE : STORE INTO VAR AS THIER ORIGINAL DATA STRUCTURE FORMAT
	// var to store json data
	var users1 user
	if checkValid {
		fmt.Println("Json data is valid")
		json.Unmarshal(jsonDataFromAPI, &users1) // we need to pass address / reference of original var
		fmt.Printf(" users1 (storing / converting JSON data into struct/ thier original data-structure format ) => %#v\n", users1)

	} else {
		fmt.Println("JSON is not valid")
	}

	//* 2ND CASE : STORING INTO MAP / KEY-VALUE PAIR FORMAT

	// var
	var user2 map[string]interface{} // interface{} we use when we will not sure of coming data's data-format

	if checkValid {
		fmt.Println("Json data is valid")
		json.Unmarshal(jsonDataFromAPI, &user2) // we need to pass address / reference of original var
		fmt.Printf(" user2 (storing / converting value into map / key-value pairs ) => %#v\n", user2)

		// printing into key-valur format
		for k, v := range user2 {
			fmt.Printf(" %v : %v ( type = %T ) \n", k, v, v)
		}
	} else {
		fmt.Println("JSON is not valid")
	}

}
