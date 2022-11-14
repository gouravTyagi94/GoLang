package main

import "fmt"

func main() {
	loginCount := 20

	if loginCount < 10 {
		fmt.Println("Number is less than 10")
	}else if(loginCount==10){
		fmt.Println("Number is equal to 10")
	}else{
		fmt.Println("Number is greater than 10")
	}


	//Special syntax used in web api's
	if num:=3; num<10{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is not less than 10")
	}
}