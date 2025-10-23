package main

import (
	"fmt"
	"time"
)

func main() {

	//* Normal code
	// greeter("Hello")
	// greeter("World")

	//* Adding goroutine to 1st function, but not giving free timw in between program execution ==> so it not get time to execute
	go greeter("Cat")
	greeter("Dog")

}

//* goroutine with time.Sleep - for  basic understanding

// Normal function without waiting in between execution
//? why waiting

// func greeter(s string) {

// 	for i := 1; i < 6; i++ {
// 		fmt.Println(s)
// 	}

// }

// * Function with time waiting in between execution
func greeter(s string) {

	for i := 1; i < 6; i++ {
		time.Sleep(5 * time.Millisecond) // This create free time of 5 milisecond in between function exceution
		fmt.Println(s)
	}

}
