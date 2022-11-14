package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)

	//Adding items in map
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	//getting a single value
	fmt.Println("JS shorts for", languages["JS"])

	//deleting an item
	delete(languages, "RB")
	fmt.Println(languages)

	//iterating over a map
	//This uses comma ok syntax(which is provided by walrus operator)
	//Therefore, we can omit either key or value
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
