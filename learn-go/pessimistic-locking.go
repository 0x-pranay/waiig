package main

import (
	"fmt"
	"sync"
)

/*

Writing concurrent program is easy.
writing correct concurrent program is difficult.
writing performant & correct concurrent program is very difficult.

Lets run 1 million threads and all of them does count++, we expect final value to be 1 million.

Since locking drastically effects performance, make sure to wrap least amount of critical code under mutex locks and keep the critical code bare minimum.

Example from Arpit Bayani - https://www.youtube.com/watch?v=4F-WiPFrPsA
*/

var count int = 0
var mu sync.Mutex

var wg sync.WaitGroup

// multiple threads can override this var
func incCount() {
	count++
	wg.Done() // done will decrement by 1
}

func doCount() {
	for i := 0; i < 1000000; i++ {
		wg.Add(1) // increments wg by 1
		go incCount()
	}
}

// Wrapping the count variable with lock and can be accessed by one thread at a time.
// Until then other threads wait for the variable to be released.
func incMutexCount() {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func doMutexCount() {
	for i := 0; i < 1000000; i++ {
		wg.Add(1) // increments wg by 1
		go incMutexCount()
	}
}

func main() {
	doCount()
	wg.Wait() //  Will wait until the final value is zero

	fmt.Printf("Without mutex: %d \n", count)

	count = 0
	doMutexCount()
	wg.Wait()
	fmt.Printf("With mutex: %d", count)
}
