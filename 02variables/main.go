package main

import "fmt"

//In Golang, starting the name of a variable is equivalent to
//Public keyword in other languages.
//If we name this as loginToken, it cant be accessed inside main func
const LoginToken = "dasdhajsdakush533"

func main() {
	var username string = "Gourav"
	fmt.Printf("Variable is of type: %T \n", username)

	var isTrue bool = true
	fmt.Printf("Variable is of type: %T \n", isTrue)

	var smallVal uint8 = 255
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.13515343843434313484
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	fmt.Println(smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	//we can declare variable without specifying the type but we can not change the
	//data type its gonna hold later in the code
	var website = "google.com"
	fmt.Println(website)
	// website = 3 This is invalid

	//No var style
	//This is done using Walrus operator
	//We can not use walrus operator outside any method, not even outside main function
	numberOfUsers := 50000
	fmt.Println(numberOfUsers)

	// numberOfUsers:= "we believe in god" This is invalid even with walrus

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
