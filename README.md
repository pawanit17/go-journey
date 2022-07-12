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
