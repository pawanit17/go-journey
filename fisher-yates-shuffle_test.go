package main

import "testing"

// Simple test driver for the fisher-yates shuffle.
func TestShuffle(t *testing.T) {

	numbers := numberlist{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers.shuffle()

	if len(numbers) != 10 {
		t.Errorf("Expected deck length of 16 but got %v", len(numbers))
	}
}
