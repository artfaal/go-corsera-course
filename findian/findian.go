package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var strInput string

	fmt.Println("Please enter a string: ")

	// Oh god! String with spaces it's something.. =)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strInput = scanner.Text()
	}

	converted_input := strings.ToLower(strInput)

	if strings.HasPrefix(converted_input, "i") && strings.Contains(converted_input, "a") && strings.HasSuffix(converted_input, "n") {
		fmt.Println("Match!")
	} else {
		fmt.Println("Not Found!")
	}
}
