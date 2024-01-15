package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	const initSize int = 3
	slice := make([]int, initSize)

	var s string
	var n int
	var err error
	count := 0

	for {
		fmt.Print("Enter integer: ")
		_, _ = fmt.Scanln(&s)
		if s == "X" {
			break
		}

		n, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error parsing integer")
			continue
		}

		if count >= initSize {
			slice = append(slice, n)
		} else { //filing in the existing slice
			fmt.Printf("Наркомаааан!")
			index := sort.SearchInts(slice, 0)
			fmt.Printf("index = %d\n", index)
			slice[index] = n
			count = count + 1
		}

		sort.Ints(slice)
		fmt.Println(slice)
	}

}
