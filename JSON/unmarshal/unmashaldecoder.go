package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	received := `{"Width" :800, "Height" : 600, "Title":"This is a image",
	"Thumbnail" : {"Url" : "http://www.example.com/image/481989943", "Height":125, "Width":100}, "Animated": false, "IDs":[116,942,234,23453]}`

	err := json.Unmarshal([]byte(received), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URL)
}
