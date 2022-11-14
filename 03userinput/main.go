package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our app: ")

	//comma ok || comma error syntax
	//This _ means error
	
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", input)
	fmt.Printf("Type of this rating is %T", input)
}