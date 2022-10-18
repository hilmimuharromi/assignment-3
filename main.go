package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	http.HandleFunc("/", StatusController)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func StatusController(w http.ResponseWriter, r *http.Request) {
	var t, err = template.ParseFiles("view.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	updateFile()
	data := readFile()
	var statusWind string
	var statusWater string

	if data.Water <= 5 {
		statusWater = "aman"
	} else if data.Water >= 6 && data.Water <= 8 {
		statusWater = "siaga"
	} else if data.Water > 8 {
		statusWater = "bahaya"
	}

	if data.Wind <= 6 {
		statusWind = "aman"
	} else if data.Wind >= 7 && data.Wind <= 15 {
		statusWind = "siaga"
	} else if data.Wind >= 15 {
		statusWind = "bahaya"
	}

	var res = map[string]interface{}{
		"Wind":        data.Wind,
		"Water":       data.Water,
		"StatusWind":  statusWind,
		"StatusWater": statusWater,
	}

	t.Execute(w, res)
}

func updateFile() {
	status := Status{}
	status.Water = rand.Intn(100)
	status.Wind = rand.Intn(100)

	content, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("status.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() Status {
	content, err := ioutil.ReadFile("status.json")
	if err != nil {
		log.Fatal(err)
	}
	status := Status{}
	err = json.Unmarshal(content, &status)
	if err != nil {
		log.Fatal(err)
	}
	return status
}
