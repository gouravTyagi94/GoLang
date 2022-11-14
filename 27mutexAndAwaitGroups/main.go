//Use command "go run --race ." to check for race condition in the program
//exit code will be 66 in this case
//mutex is used to overcome this race condition

// Why do we need RWMutex?
// When using Mutex the value from the memory will be locked
// until the Unlock method will be invoked.
// This is also valid for the reading phase. In order to make reading accessible for multiple threads,
// the Mutex can be replaced with RWMutex, and for reading it will be used RLock and RUnlock methods.

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions")

	//This time we are using waitgroup as pointer variable
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	//we can also use mut := &sync.RWMutex here, it would
	//change nothing in the output here

	var score = []int{}
	//As we are adding 3 goroutines, we can also use
	//wg.Add(1) three times
	wg.Add(3)

	//These are immediately invoked functions
	//These are anonymous and are invoked just after their creation
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
