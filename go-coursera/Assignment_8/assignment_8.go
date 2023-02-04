package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	// store created animals <name, obj>
	existAniMap := make(map[string]Animal)

	// main procedure
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println(">")

		// read input
		scanner.Scan()
		inputLine := scanner.Text()
		fields := strings.Fields(inputLine)

		// dispatch according to command
		switch {
		case fields[0] == "newanimal":
			animal := createAnimal(fields[1], fields[2])
			existAniMap[fields[2]] = animal

		case fields[0] == "query":
			queryAnimal(existAniMap, fields[1], fields[2])
		}
	}
}

// createAnimal is for creating new animal
func createAnimal(category string, uniName string) Animal {
	var ani Animal
	switch {
	case category == "cow":
		ani = Cow{uniName}
	case category == "bird":
		ani = Bird{uniName}
	default:
		ani = Snake{uniName}
	}
	fmt.Println("Created it!")
	return ani
}

// queryAnimal is for executing query animal's action info
func queryAnimal(existAnimals map[string]Animal, uniName string, action string) {
	animal := existAnimals[uniName]
	if nil != animal {
		if "eat" == action {
			animal.Eat()
		} else if "move" == action {
			animal.Move()
		} else {
			animal.Speak()
		}
	}
}
