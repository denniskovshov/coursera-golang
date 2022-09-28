package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := GetInputNumbers()

	BubbleSort(numbers)

	fmt.Println(numbers)
}

// not optimized version
func BubbleSort(numbers []int) {
	sorted := false

	for !sorted {
		// assume it's sorted, code inside next loop will disable that if needed
		sorted = true

		for i := 0; i < len(numbers)-1; i++ {
			// compare and swap
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)

				// if array is fully sorted this will never trigger
				// so 'sorted' will stay 'false' and loop will end
				sorted = false
			}
		}
	}

}

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}

func GetInputNumbers() []int {
	// collect user input
	fmt.Println("Enter up to 10 numbers separated by space:")

	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	input := inputScanner.Text()

	// expecting numbers separated by space
	numberStrings := strings.Split(input, " ")

	numbers := make([]int, 0, 10)

	// convert strings to numbers
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}

	return numbers
}
