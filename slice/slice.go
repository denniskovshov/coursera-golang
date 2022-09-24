package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		// collect user input
		fmt.Println("Please, enter an integer or X to exit:")

		var input string
		fmt.Scan(&input)

		if input == "X" {
			fmt.Println("Exiting...")
			break
		}

		// convert to integer
		num, err := strconv.Atoi(input)

		if err != nil {
			continue
		}

		// add to slice, sort, print
		slice = append(slice, num)

		sort.Sort(sort.IntSlice(slice))

		fmt.Println(slice)
	}
}
