package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Banana"}

	fruitList = append(fruitList, "Mango", "Peach")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	//Declaring slice using make() function
	highScores := make([]int, 4)
	highScores[0] = 1531
	highScores[1] = 6564
	highScores[2] = 4848
	highScores[3] = 6651
	//highScores[4] = 5646;

	//This will throw error as highScores[4] is out of range
	// fmt.Println(highScores)

	//But this will work totally fine as append function
	//reallocates the memory
	highScores = append(highScores, 23231, 12121, 5485)
	fmt.Println(highScores)

	//checking whether int slice is sorted or not
	fmt.Println(sort.IntsAreSorted(highScores))

	//Sorting an int slice
	sort.Ints(highScores)
	fmt.Println(highScores)


	//How to remove a value from a slice using append function

	var courses = []string{"JS", "Java", "C++", "Python", "Golang"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
