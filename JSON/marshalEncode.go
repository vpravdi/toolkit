package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname, Lname string
	Items        []string
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encoder)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>BOO</title>
			</head>
			<body>
			You are at booooo
			</body>
			</html>`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	p1 := person{
		"Dinesh",
		"Praveen",
		[]string{"a", "b", "c"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encoder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	p1 := person{
		"Dinesh",
		"Praveen",
		[]string{"a", "b", "c"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}

}
