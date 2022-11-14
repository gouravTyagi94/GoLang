package main

import "fmt"

func main() {

	//We can not create functions inside main function in golang
	greeter()

	//Simple adder
	result := adder(3, 5)
	fmt.Println("Result is ", result)

	//Proadder
	proResult := proAdder(3, 5, 15)
	fmt.Println("Pro Result is ", proResult)

	//Proadder with 2 return values
	proAdd2, myMessage := proAdder2(2, 3, 5, 6)
	fmt.Println(proAdd2, myMessage)
}

func adder(a int, b int) int {
	return a + b
}

//These are variadic functions in golang
//Similar to varargs in Java
func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}

//function with multiple return values
func proAdder2(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "Hello World!!!"
}

func greeter() {
	fmt.Println("Hello world!!!")
}
