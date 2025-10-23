package main

import (
	"fmt"
	"net/http"
	"sync"
)

//! In this examples its hard to notice error in Shared Memory , The use of Mutex Lock() & Unlock() .. but in real life projects the differenc is very much

var wg sync.WaitGroup

//! In this we actully pass *pointer, here are not doing that to keep things simple

// This manages to pick a time for gourtines to exceute between in process we get time
// like time we are creating using time.Sleep fucntion

// so as you know in real life task request & responses , when programe get free time , for e.g. time get while waiting from data to come from DB that is in another country
// so using that time efficently is handle by this sync package

// type WaitGroup   <= special type , struct

// and these 3 methods

// func (wg *WaitGroup) Add(delta int)
// func (wg *WaitGroup) Done()
// func (wg *WaitGroup) Wait()

// * One more imp sync.Mutex() Method
var mutex sync.Mutex

//! In this we actully pass *pointer, here are not doing that to keep things simple
/*
* 🧠 Concept

A sync.Mutex (mutual exclusion lock) is used to prevent multiple goroutines from accessing shared data at the same time.
It ensures that only one goroutine can enter a critical section (a code block that modifies shared state) at once.

Without a mutex → you’ll get race conditions, meaning your program behaves unpredictably due to concurrent access.


* 📘 Why sync.Mutex is Necessary

🧩 Avoids race conditions: ensures consistent data updates.
🔒 Protects shared memory: when multiple goroutines read/write same variable.
⚡ Performance: simpler and faster than channels for just locking data access.

🚫 Common Mistakes

1. Forgetting to unlock → deadlock (program freezes).
2. Locking around code that doesn’t need protection → unnecessary slowdown.
3. Unlocking without locking → panic.

*/

// * Shared memory ( critical section )
var signals = []string{"Test"}

func main() {

	// slice
	websitesList := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",
		"https://amazon.in",
	}

	for _, web := range websitesList {

		go SendStatusCode(web)
		//This is main built in Go feature , This concurrenty nature is not given by that sync package , That just manage when timing, When goroutines go , come & done , and wait for all of them to come
		// Now it in goroutine , It handle itself that "Concurrent Nature", Means get execute when get free time

		wg.Add(1)
		//As i said we can image this package make List of goroutines
		// So this is the method which add goroutines in to that list, So this Sync package can manage time,kepp track of all goroutines & wait for it
	}

	// This methods work is stop main function from completing , It wait until the goroutines fully get time and complete their execution and get back reult or specially say until they send "Done" signel
	wg.Wait() // basically you can image, This package create List of all goroutines, and this methods wait for all do finish thier work and send Done Signal

	fmt.Println("")
	fmt.Println(signals)

}

// Function that send GET request and send StatusCode as response

func SendStatusCode(website string) {

	defer wg.Done() // This method send signal to that Sync package so it can keep a track of which goroutines is Done , or still in execution need more time
	// So after compeletion of Task , That goroutines carry , it send Done Signal

	req, err := http.Get(website)

	if err != nil {
		fmt.Printf("\nOOps got err in %s \n", website)
	} else {

		//* Image this as a Special Room , And 2 gates, one for In and other for Out

		mutex.Lock() // Lock: Only one goroutine can enter from this point , and access below code

		signals = append(signals, website) // Shared Memory ( Critical Section )

		mutex.Unlock() // Unlock: After Completion of first entered goroutine's work , it allow next waiting goroutine to enter in this critincal section

		fmt.Printf("\nSuccessfull, Status Code for %s is %d\n", website, req.StatusCode)
	}

}

/*

* Without goroutines and sync
red-dragon@red-dragon-HP-Laptop-15s-fq4xxx:/mnt/Dev/backend-learning/0.8.2-goroutines-sync$ go run main.go

Successfull, Status Code for https://google.com is 200

Successfull, Status Code for https://fb.com is 200

Successfull, Status Code for https://github.com is 200

Successfull, Status Code for https://go.dev is 200

Successfull, Status Code for https://amazon.in is 200


* With goroutines and sync
red-dragon@red-dragon-HP-Laptop-15s-fq4xxx:/mnt/Dev/backend-learning/0.8.2-goroutines-sync$ go run main.go

Successfull, Status Code for https://github.com is 200

Successfull, Status Code for https://google.com is 200

Successfull, Status Code for https://go.dev is 200

Successfull, Status Code for https://fb.com is 200

Successfull, Status Code for https://amazon.in is 200



*/
