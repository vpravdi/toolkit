package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	a := "i am a rockstar go developer"

	encodedString := base64.StdEncoding.EncodeToString([]byte(a))

	fmt.Println(encodedString)

	decodedString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decodedString))
}
