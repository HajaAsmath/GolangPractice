package main

import (
	"fmt"
	greetings "greetingspack"
	"log"
)

func main() {
	log.SetPrefix("Greeting:")
	message, err := greetings.Greetings("Haja")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
