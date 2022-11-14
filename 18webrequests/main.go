package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	response, err := http.Get(url)

	if err!=nil{
		panic(err)
	}
	fmt.Printf("Response is of type: %T \n", response)

	defer response.Body.Close() //Caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err!=nil{
		panic(err)
	}
	fmt.Println(string(databytes))
}