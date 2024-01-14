package main

import "fmt"

func main() {
	fmt.Println("Please enter a floating point number: ")
	var input float64
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid floating point number.")
		return
	}
	output := int(input)
	fmt.Printf("Your number is %d", output)
}
