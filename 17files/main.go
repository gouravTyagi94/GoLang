package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Working with files in Golang"
	file, err := os.Create("./file.txt")

	// if err!=nil{
	// 	panic(err)
	// }
	// using check error function in place of if
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	
	checkNilErr(err)

	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./file.txt")
}

func readFile(filename string){
	//The file that we are reading is never read in string format.
	//Instead, it is read in byte form. Normally, files on internet
	//JSON data is found in byte form.
	databyte, err := ioutil.ReadFile(filename)
	
	checkNilErr(err)

	//output will be in bytes
	// fmt.Println("The data inside file is: \n", databyte)

	//Convert byte output into string using typecasting
	fmt.Println("The data inside file is: \n", string(databyte))
}


func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}