package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// collect user input
	var input string
	fmt.Println("Please enter a string:")

	// bufio is needed to collect entire string even with spaces
	console_reader := bufio.NewScanner(os.Stdin)

	console_reader.Scan()
	input = console_reader.Text()

	// normalize for comparison
	input = strings.ToLower(input)

	// check
	startsWithI := strings.HasPrefix(input, "i")
	containsA := strings.Contains(input, "a")
	endsWithN := strings.HasSuffix(input, "n")

	// final output
	if startsWithI && containsA && endsWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
