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
	// Tags     []string `json:"tags, omitempty"` //* this "omitempty is exclude filed when value is nil"
	Tags []string `json:"tags,omitempty"` // this "omitempty is exclude filed when value is nil" //! carefully write synatx, dont give space
}

func main() {

	fmt.Println("Welcome to Mashelling Tutorial of Go")
	// EncodeJson()
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

	// package this data as json data

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

// Decode Json , for Data come from API suppose and want to consume as JSON

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

	// var to store json data
	var users user
	if checkValid {
		fmt.Println("Json data is valid")
		json.Unmarshal(jsonDataFromAPI, &users) // we need to pass address of original var
		fmt.Printf("%#v\n", users)

	} else {
		fmt.Println("JSON isn ot valid")
	}
}
