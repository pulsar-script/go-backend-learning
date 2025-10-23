package main

import (
	"fmt"
	"sync"
)

//* Use "--race" flag to check race condition , it is a tool
// e.g. go run --race .       (. means every file)
// e.g. go run --race main.go

// Creating sync.Waitgroup & sync.Mutex type vars using proper pointers

func main() {

	var wg = &sync.WaitGroup{}
	var mutex = &sync.Mutex{}

	// Read-Write Mutex , it work exatcly same as Mutex for write , additionally it have Read-Lock RLock & Runlock
	var RWmutex = &sync.RWMutex{}

	// score
	score := []int{0}

	fmt.Println("Welcome to go-lang Race Condtion Tutorial")

	// add into list
	wg.Add(4) // you can individually add before each func , but this is more common way - no side  effects

	// These are "INF" - Imediatly Invoke Function
	// They are not much special
	func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()
		fmt.Println("INF 1")

		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()

	}(wg, mutex)

	func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()
		fmt.Println("INF 2")

		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()

	}(wg, mutex)

	func(wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()
		fmt.Println("INF 3")

		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()

	}(wg, mutex)

	func(wg *sync.WaitGroup, RWmutex *sync.RWMutex) {

		defer wg.Done()
		fmt.Println("INF 4")

		RWmutex.RLock() //* GOOD PRACTICE : The main puropse is its , When someown is writing , That time if someone try to read it , It not allow or remove from Reading Permission
		fmt.Printf("Score Reading in INF 4 %v", score)
		RWmutex.RUnlock()

	}(wg, RWmutex)

	wg.Wait()

	fmt.Println("")
	fmt.Println(score)

}
