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
- Functions in GO can return more than one value as result.
```
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Reached the main method")

	sqrtRoot, problem := sqrt(-32)

	if problem != nil {
		fmt.Print(problem)
	} else {
		fmt.Print(sqrtRoot)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	} else {
		return math.Sqrt(x), nil
	}
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

## Maps
- Maps map from a key to a value.
- Map is created using the built-in ```make``` method.
- An entry in the map is deleted via the ```remove``` method.
- A value in an entry in a Map can be accessed via the [].
```func main() {
	fmt.Println("Reached the main method")

	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["quadrilateral"] = 4
	vertices["pentagon"] = 5
	fmt.Println("Reached the called method!", vertices)

	fmt.Println("Reached the called method!", vertices["pentagon"])

	delete(vertices, "pentagon")

	fmt.Println("Reached the called method!", vertices["pentagon"])
}
```

## Loops
- Only the for loop is available in GO.
- ```
func main() {
	fmt.Println("Reached the main method")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

- Using ```Range```
```
func main() {
	fmt.Println("Reached the main method")

        // Range on a string
	names := [4]string{"Naruto", "Sasuke", "Ichigo", "Bankai"}

	for index, value := range names {

		fmt.Print(index)
		fmt.Println(value)
	}
	
	// Range on a Map
	vertices := make(map[string]int)
	vertices["alpha"] = 1
	vertices["beta"] = 2
	vertices["charlie"] = 3
	vertices["delta"] = 4
	vertices["echo"] = 5
	vertices["foxtrot"] = 6

	for key, value := range vertices {

		fmt.Print(key)
		fmt.Println(value)
	}
}


