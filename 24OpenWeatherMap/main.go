package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to OpenWeather Map services")
	r := mux.NewRouter()
	r.HandleFunc("/{city}", getWeatherData)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func getWeatherData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	query := params["city"]
	apiKey := "81e499831717f237a109d9fef6546d81"
	units := "metric"
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + query + "&appid=" + apiKey + "&units=" + units)

	// var ans string
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var ans interface{}
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(body, &ans)

	json.NewEncoder(w).Encode(ans)
}
