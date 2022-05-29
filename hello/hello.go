package main

import (
	"example.com/greetings"
	"fmt"
	"golang.org/x/example/stringutil"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	// test for calling with another api
	fmt.Println(stringutil.Reverse("Hello"))
}
