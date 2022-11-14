package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	//using aliases in struct
	//These will be useful when we convert struct into json

	//Here, Name will be shown as coursename when we convert struct into json
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`

	//This - shows that we dont want to show this field to users of our API's
	Password string `json:"-"`

	//This means that we want to show Tags field as tags and if the field is nil,
	//omit the tag. But we need to make sure that there is no space in between , and omitempty,
	//Otherwise, this will show error
	Tags []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abcd1234", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abcdsd1234", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 199, "LearnCodeOnline.in", "abdawcd1234", nil},
	}

	//package this as json data

	//If we use Marshal method, data received as json would be in unreadable format
	//so its better to use MarshalIndent

	// finalJson, err := json.Marshal(courses)

	//2nd param is prefix, and 3rd is type of indent
	//Run and check the output to visualise
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}
	//if we use println here, data would be print in bytes
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "Angular Bootcamp",
			"Price": 199,
			"website": "LearnCodeOnline.in"
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was Valid")

		//Unmarshalling the marshalled data
		//We are passing lcoCourse as reference because we want the unmarshalled data to be saved in
		//lcoCourse variable
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		//To print this type of data, we have a special specifier
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//Some cases where u just want to add data to key value

	//Here, we are asssuming that we are fetching data from web and creating a map
	//The key is guaranteed to be a string but value can be of any type, so we are
	//assuming value to be of interface (struct in golang) type.
	//Doing this will accomodate all data types
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v, value is %v and type is %T\n", k, v, v)
	}
}
