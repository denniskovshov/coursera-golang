package main

import "fmt"

func main() {
	var num_input float64
	var num_result int

	// collect user input
	fmt.Println("Please enter a floating point number:")
	_, err := fmt.Scan(&num_input)

	if err != nil {
		fmt.Println(err)
		return
	}

	// convert to int and output
	num_result = int(num_input)
	fmt.Printf("Integer is: %d \n", num_result)
}
