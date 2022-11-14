package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, err:= ioutil.ReadAll(response.Body);

	//Brute force way of typecasting data inside "content" which is in byte format
	//into string format
	fmt.Println(string(content))
	
	//Another method is to use strings library which provides a stringbuilder
	var responseString strings.Builder
	
	//bytelength gives the len of content, it is similar as response.ContentLength .

	// using Write method to write the byte data of "content" into response string
	//Now, this responseString contains the content in string format 
	 
	byteCount, err:= responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println("Content is ", responseString.String())	
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"courseName" : "Let's go with Golang",
		"price" : "0",
		"platform" : "Youtube"
	}
	`)

	//Sending the post request
	response, err :=http.Post(myurl, "application/json", requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	//We need to create form data this time
	//Whenever we post any data, the data can be accessed using url package
	//So, url package will give us access to the data that we just posted

	data := url.Values{}
	data.Add("First Name", "Gourav")
	data.Add("Email", "Gourav@go.dev")
	response, err :=http.PostForm(myurl, data)

	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _:= ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}