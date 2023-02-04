package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

// initAnimal is for initializing a specific animal
func (a *Animal) initAnimal(food, loco, sound string) {
	a.food = food
	a.locomotion = loco
	a.noise = sound
}

// getAnimalInfo is for returning related method of an animal
func (a *Animal) getAnimalInfo(method string) func() string {
	if "eat" == method {
		return a.Eat
	} else if "move" == method {
		return a.Move
	} else {
		return a.Speak
	}
}

func main() {

	// initialization
	var cow, bird, snake Animal
	cow.initAnimal("grass", "walk", "moo")
	bird.initAnimal("worms", "fly", "peep")
	snake.initAnimal("mice", "slither", "hsss")

	animalMap := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	// main procedure
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println(">")

		// read input
		scanner.Scan()
		inputLine := scanner.Text()
		fields := strings.Fields(inputLine)

		// parse & calling
		inputAnimal := animalMap[fields[0]]
		animalInfo := inputAnimal.getAnimalInfo(fields[1])
		fmt.Println(animalInfo())
	}
}
