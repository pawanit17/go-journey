# Introduction
- Strongly typed language
- Has its own Garbage collection scheme
- Inherently parallel processing language
- Developed at Google
- 12 minute video coverage of GO.

# What problem does GO solve
- Go was developed in Google when the developers noticed that the fastness of C/C++ is offset by the slower build processes.
- They also wanted to have a simpler set of language constructs - simplicity.

# Samples
- main method inside main package is the entry point into the application.
- The return type is specified at the end of the method signature unlike other languages.
- Default value is the zero value of the data type - 0 for int, empty string for string.

## Variables & Functions 
```
package main

import "fmt"

func main() {
	fmt.Println("Reached the main method")

	var name string = "Linku"
	var index int = 3
	value := count(name, index)

	fmt.Println("value:", value)
}

func count(teddyName string, index int) int {
	fmt.Println("Reached the called method!", teddyName)
	fmt.Println("Reached the called method!", index)

	return -1
}
```

## Arrays & Slices
- By default, all elements in the array are initialized to 0.
```
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)

	numbers[2] = -1
	fmt.Println(numbers)
```
- Array size has to be determined at compile time. So we cannot add a sixth element to a 5 element array.
- Also the size of the array cannot be dynamically provided. It should be available as a number when the program is compiling.
- Slices are used to overcome this. Slices are abstraction over array to simulate dynamic number of elements.
- Append method can be used on Slice to update it. Slices are backed up by Arrays.
- ```func main() {
	fmt.Println("Reached the main method")

	numbersArray := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbersArray)

	numbersSlice := []int{10, 20, 30, 40, 50}
	fmt.Println(numbersSlice)

	numbersSliceUpdated := append(numbersSlice, 60)
	fmt.Println(numbersSlice)
	fmt.Println(numbersSliceUpdated)
     }```
Gives:
```
PS D:\Development\Go\Basics> go run basics.go
Reached the main method
[10 20 30 40 50]
[10 20 30 40 50]
[10 20 30 40 50]
[10 20 30 40 50 60]
```


