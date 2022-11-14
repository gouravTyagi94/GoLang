package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
		//if we mention a fallthrough, the next case will also be executed.
		
		//Fall through is a type of error that occurs in various programming languages
		//like C, C++, Java, Dart …etc. It occurs in switch-case statements where when we
		//forget to add a break statement and in that case flow of control jumps to the next line.

		//We dont need to use break in golang, it automatically handles this
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again")
	}
}
