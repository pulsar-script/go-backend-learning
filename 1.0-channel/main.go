//* To run steps and see code's outpus ,properly uncomment that step code and comment down other step code , they cause error

package main

import (
	"fmt"
	"sync"
)

//* In this code we just learn channels implementations & basic , To brush up watch Hitesh's video
//? Still not understand its purpose , benefits

func main() {

	fmt.Println("Welcome to Channels tutorial")

	// Normal channel declaration
	// myCh := make(chan int)

	// Decleration of Buffered Channel
	myCh := make(chan int, 4) // it work with looping thorugh single listing channel statement , 2 mean  i want to add 2 more value
	//* Read more about it

	//* STEP : 1 LEARNING
	// how to push value in Channel
	// myCh <- 5 // we push 5 value in our channel

	// lets try to receive value from channel or say listing to channel

	// fmt.Println(<-myCh) //! but it wont work

	//* Output get when tyr to run

	/*

		Welcome to Channels tutorial
		! fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
		        /mnt/Dev/backend-learning/1.0-channel/main.go:15 +0x6e
		exit status 2

	*/

	//? why it wont work ?
	// => beacuse imagine channels as piplines, when we push value that time someone should be listing to channel then only channels work
	// But, yaa we can create channel

	// SOLUTION => create goroutines

	//* STEP : 2 LEARNING

	wg := &sync.WaitGroup{}

	wg.Add(4)

	// 1st goroutine for listing channel
	go func(myCh chan int, wg *sync.WaitGroup) {

		defer wg.Done()
		// fmt.Println(<-myCh)

	}(myCh, wg)

	// 2nd goroutine for pushing value in channel
	go func(myCh chan int, wg *sync.WaitGroup) {

		defer wg.Done()
		// myCh <- 5

		//! seconding 2nd value ,but their is only one listner
		// myCh <- 6  //! it wont work cause error , their should be equal number of senders and listerns ( we can loop through single listing statement upto listners number of time  )

	}(myCh, wg)

	//* OUTPUT FOR STEP 2

	/*

	   red-dragon@red-dragon-HP-Laptop-15s-fq4xxx:/mnt/Dev/backend-learning/0.10-channel$ go run main.go
	   Welcome to Channels tutorial
	   5

	*/

	//* STEP : 3 LEARNING

	//* we can close channel also , mostle we do closing in pushing section

	//* Also insted of using only simple chan type , we specifiy channels direction
	// is it listing channel "<-chan"  or pushing channel "chan<-"

	// LISTIEN ONLY CHANNEL
	go func(myCh <-chan int, wg *sync.WaitGroup) {

		defer wg.Done()

		// To check properly that 0 is closed Channel signal or pushed value 0

		val, isChannelOpen := <-myCh

		if isChannelOpen {

			fmt.Printf("Channel is open and the value pushed is %d", val)

		} else {

			fmt.Println("Channel is Closed")
		}

		fmt.Println("")
		// fmt.Println(<-myCh)

	}(myCh, wg)

	// SEND ONLY CHANNEL
	go func(myCh chan<- int, wg *sync.WaitGroup) {

		defer wg.Done()

		// close(myCh) //! It throw error , Dont close channel and after that try to push values
		// myCh <- 5
		// myCh <- 6

		//! but when we close channel without pushing any value it send back zero

		//? But why if someone send zero as value , how to identify which is close signal and which is pushed value 0
		myCh <- 0

		close(myCh)

	}(myCh, wg)

	wg.Wait()
}
