package main

import (
	"fmt"
)

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
	for true {
		fmt.Println(">")

		// read input
		var name, action string
		_, err := fmt.Scanf("%s %s", &name, &action)
		if err != nil {
			fmt.Println("error")
			return
		}

		// parse & calling
		inputAnimal := animalMap[name]
		animalInfo := inputAnimal.getAnimalInfo(action)
		fmt.Println(animalInfo())
	}
}
