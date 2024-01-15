package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// I read the assignment several times. And as it seems to me quite contradictory.
	// They say an empty slice, BUT with a length of three. It's no longer empty.
	// So here are two implementations. I really don't know which one is correct,
	// but I chose the one that makes more sense for me.

	// batch := make([]int, 3)

	var input string
	errorMsg := "Invalid input. Please enter a valid int number."
	batch := make([]int, 0, 3)
	fmt.Println("Hello! For exit enter \"X\"")

	for {
		fmt.Println("Please insert integer number")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println(errorMsg)
			continue
		}
		// Found this cool strings.EqualFold function. Ignore case when comparing strings =)
		if strings.EqualFold(input, "X") {
			fmt.Println("Goodbye!")
			break
		}
		num, err := strconv.Atoi(input) // Convert to int
		if err != nil {
			fmt.Println(errorMsg)
			continue
		}
		batch = append(batch, num)
		sort.Ints(batch)
		fmt.Println("Current batch:", batch)
	}
}
