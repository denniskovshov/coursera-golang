package main

import (
	"fmt"
)

func main() {
	for {
		// collect user input
		fmt.Println("Enter animal (cow,bird,snake) and action (eat,move,speak):")
		var animalInput, actionInput string
		fmt.Scan(&animalInput, &actionInput)

		// create animal object from input
		animal := createAnimal(animalInput)

		if animal == (Animal{}) {
			fmt.Printf("Unknown animal: %v\n", animalInput)
			continue
		}

		// get animal animalAction from input
		animalAction := getAnimalAction(animal, actionInput)

		if animalAction == nil {
			fmt.Printf("Unknown action: %v\n", actionInput)
			continue
		}

		// final output
		animalAction()
	}
}

type Animal struct {
	food   string
	motion string
	noise  string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.motion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func createAnimal(animalInput string) Animal {
	var animal Animal

	switch animalInput {
	case "cow":
		animal = Animal{"grass", "walk", "moo"}

	case "bird":
		animal = Animal{"worms", "fly", "peep"}

	case "snake":
		animal = Animal{"mice", "slither", "hsss"}
	}

	return animal
}

func getAnimalAction(a Animal, actionInput string) func() {
	var action func()

	switch actionInput {
	case "eat":
		action = a.Eat

	case "move":
		action = a.Move

	case "speak":
		action = a.Speak
	}

	return action
}
