package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	// collect user input
	fmt.Println("Enter file name:")

	var fileName string
	fmt.Scan(&fileName)

	// open file
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// read file line by line
	names := make([]Name, 0, 100)

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		// split line
		textLine := fileScanner.Text()
		textParts := strings.Split(textLine, " ")

		// create object
		firstName := textParts[0]
		lastName := textParts[1]
		nameFromFile := Name{fname: firstName, lname: lastName}

		// add
		names = append(names, nameFromFile)
	}

	// output
	fmt.Println()

	for _, name := range names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}
