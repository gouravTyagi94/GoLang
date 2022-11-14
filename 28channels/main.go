//Channels are a way in which multiple goroutines can talk
//to each other. One go routine can not have information regarding
//the expected time required for completing the task by another goroutine
//or whats going inside it, but it may be the case that one go routine is
//just requiring a signal from another to do some work.


//Channel are by default bidirectional in nature but we can specify a 
//channel to be send only or receive only channel using this syntax

// Receive only
// go func(ch <-chan int, wg *sync.WaitGroup)

//Send Only
// go func(ch chan<- int, wg *sync.WaitGroup)
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	//no of listeners should be equal to no of values passed in channel
	//otherwise error will occur. To overcome this, we can use buffered channel
	//like this
	// myCh := make(chan int, 1) or myCh := make(chan int, 5)

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		//Two listeners(println statements) are required to
		//read two values from channel, otherwise it will
		//throw error
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)

		//Listening on a closed channel is valid but sending is not
		//Also, if we listen to a closed channel without sending anything
		//we will receive 0 on printing, similarly if we send 0 and print it
		//we recieve 0. So, how can we differentiate between these two conditions.

		//By receiveing value instead of printing, and then checking the value of
		//boolean isChannelOpen, we can differentiate between these two conditions
		val, isChannelOpen := <-myCh
		if isChannelOpen {
			fmt.Println(val)
		}
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh<-5
		myCh<-6
		myCh<-7
		//close the channel
		close(myCh)

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
