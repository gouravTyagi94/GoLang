package main

import "fmt"

func main() {

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)
	// for i:=0; i<len(days); i++{
	// 	fmt.Println(days[i])
	// }


	//Here i just shows the index of current element in slice
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	//For each loop in golang
	// for index, day:= range days{
	// 	fmt.Printf("The index is %v and the day is %v\n", index, day)
	// }

	rougueValue:= 1

	//This is similar to while loop
	for rougueValue<10{

		//usgae of goto statement
		if(rougueValue==2){
			goto gourav
		}
		if(rougueValue==5){
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

	gourav:
		fmt.Println("Golang is good")
}