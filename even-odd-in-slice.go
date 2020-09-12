// First Go program to iterate through a list of number in the slice
// and determine whether it is an even number of an odd number.
package main

import "fmt"

func main() {

	// Create a slice
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println("Even Number")
		} else {
			fmt.Println("Odd Number")
		}
	}
}
