package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //usually this variable is pointer
var mut sync.Mutex   //usually this variable is pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"http://www.google.com",
		"http://fb.com",
		"http://github.com",
		"http://twitter.com",
	}
	for _ , web := range websitelist{
		//simply writing below statement is not suffice because
		//we are not waiting for the goroutine to come back and report us.
		//we need to use sync.Waitgroup for this(can do this using time.sleep()
		//as well but waitgroup is better)
		//go getStatusCode(s)
		
		
		go getStatusCode(web)
		//for now, just assume that the 1 we are passing means that
		//there is only one go coroutine
		wg.Add(1)
	}


	//This wait method is usually called at the end of method(main
	//or any other method in which it is used).
	//This prevents the termination of method without succesful completion
	//of all go coroutines. 
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	//It is our responsibility to use defer keyword and mark the
	//task as done after completing all the steps in the function
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops, an error occured...")
	}else{
		//this is mutual exclusion lock
		//we are locking the storage before writing 
		//and unlock after we are done with it
	
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
