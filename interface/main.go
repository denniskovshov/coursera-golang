package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	nameAnimalMap := make(map[string]*Animal)

	for {
		input := getUserInput()

		err := validateInput(input)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if input[0] == "newanimal" {
			name := input[1]
			animalType := input[2]

			newanimal := createAnimal(animalType)
			nameAnimalMap[name] = newanimal

			fmt.Println("Created it!")
			continue

		} else if input[0] == "query" {
			name := input[1]
			actionName := input[2]

			animal, ok := nameAnimalMap[name]

			if !ok {
				fmt.Printf("'%s' name not found\n", name)
				continue
			}

			action := getAnimalAction(animal, actionName)
			action()
		}
	}
}

// "Instructions" has a typo where the interface should not contain 3 string fields
// but rather 3 methods as described in "My submission"
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c *Cow) Eat()   { fmt.Println("grass") }
func (c *Cow) Move()  { fmt.Println("walk") }
func (c *Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (c *Bird) Eat()   { fmt.Println("worms") }
func (c *Bird) Move()  { fmt.Println("fly") }
func (c *Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (c *Snake) Eat()   { fmt.Println("mice") }
func (c *Snake) Move()  { fmt.Println("slither") }
func (c *Snake) Speak() { fmt.Println("hsss ") }

func createAnimal(animalType string) *Animal {
	var animal Animal

	switch animalType {
	case "cow":
		animal = new(Cow)

	case "bird":
		animal = new(Bird)

	case "snake":
		animal = new(Snake)
	}

	return &animal
}

func getAnimalAction(a *Animal, actionName string) func() {
	var action func()

	switch actionName {
	case "eat":
		action = (*a).Eat

	case "move":
		action = (*a).Move

	case "speak":
		action = (*a).Speak
	}

	return action
}

func getUserInput() []string {
	fmt.Println("Enter command: [newanimal <name> <cow,bird,snake>, query <name> <eat,move,speak>]")

	stdinReader := bufio.NewScanner(os.Stdin)
	stdinReader.Scan()
	input := stdinReader.Text()

	inputArgs := strings.Split(input, " ")

	return inputArgs
}

func validateInput(input []string) error {
	command := input[0]

	if command != "newanimal" && command != "query" {
		return errors.New("Command should either be 'newanimal' or 'query'")
	}

	if command == "newanimal" {
		if len(input) < 3 {
			return errors.New("Incorrect 'newanimal' arguments")
		}
	}

	if command == "query" {
		if len(input) < 3 {
			return errors.New("Incorrect 'query' arguments")
		}
	}

	return nil
}
