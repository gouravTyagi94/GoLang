package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("Value of ptr is ", ptr)

	myNum := 23

	var ptr = &myNum

	fmt.Println("Address of nyNum is ", ptr)
	fmt.Println("Value to which ptr is referencing ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is ", *ptr)

}
