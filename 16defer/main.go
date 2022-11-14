package main

import "fmt"

func main() {

	//deferred lines are executed just before the end of the function
	//It works in LIFO manner, so the last deferred statement will be executed first( but normal code
	// lines are obviously executed first)
	//Output here:- Hello world, Three, Two, One
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello World!!!")
	myDefer()
}

func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}