// Package for holding the implementation for the Fisher-Yates shuffle over a slice in Go!.
// Examples:
// Run1: go run fisher-yates-shuffle.go
// [1 2 3 4 5 6 7 8 9 10]
// [9 1 2 3 4 5 6 7 10 8]

// Run 2: go run fisher-yates-shuffle.go
// [1 2 3 4 5 6 7 8 9 10]
// [3 1 10 2 4 5 6 7 8 9]
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type numberlist []int

func main() {

	numbers := numberlist{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers.print()
	numbers.shuffle()
	numbers.print()
}

// shuffle as the name suggests, shuffles the underlying slice.
// The algorithm that is used for this shuffle is Fisher-Yates.
func (l numberlist) shuffle() {

	for i := 0; i < len(l); i++ {

		// UnixNano gives an ever changing number, which can be used to serve as the Pseudo Random Number generator seed.
		newRand := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Get a random number between 0 and len-1
		j := newRand.Intn(len(l))

		// Swap the slice at 'i' and 'j' indexes.
		l[i], l[j] = l[j], l[i]
	}
}

// print prints the list of numbers so that it could be used before
// and after the shuffle to see the completeness of the shuffule.
func (l numberlist) print() {

	fmt.Println(l)
}
