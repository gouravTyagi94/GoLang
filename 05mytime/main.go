package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	//We need to use exactly this string "01-02-2006 Monday 15:04:05" to
	//get the output formatted according to our needs.

	//get all three, date, day and time
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	//get date and day only
	fmt.Println(presentTime.Format("Monday 01-02-2006"))

	//create a date
	createdDate := time.Date(2020, time.March, 6, 16, 0, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	//To build windows, mac and linux executables, commands are
	// GOOS = "windows" go build
	// GOOS = "darwin" go build
	// GOOS = "linux" go build
}
