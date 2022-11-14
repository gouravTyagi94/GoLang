package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please rate our app between 1 and 5:")
	input, _:= reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", input)

	numRating, err:= strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to the rating: ", numRating+1)
	}
}