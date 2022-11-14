package main

import "fmt"

func main() {
	var fruitList[4] string

	//len of array in golang always shows the size and not the number of items
	//in array
	fmt.Println(len(fruitList))

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"


	fmt.Println(fruitList)

	var vegList = [5]string{"potato", "tomato", "brinjal"}
	fmt.Println(len(vegList))
	fmt.Println(vegList)

}