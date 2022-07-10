# Introduction
- Strongly typed language
- Has its own Garbage collection scheme
- Inherently parallel processing language
- Developed at Google

# What problem does GO solve
- Go was developed in Google when the developers noticed that the fastness of C/C++ is offset by the slower build processes.
- They also wanted to have a simpler set of language constructs - simplicity.

# Deep dive
- Variables
  ```var i int = 42```
  ```var j float32 = 9.4
  ```fmt.Printf("%v, %T", i, i )```
- var block
  - ```
    var (
       abc = 10
       def = 9.0
       ghi ="P"
    )
    ```
- Shadowing
   - Variable that is in the innermost scope takes preference. 
  
- Expose
  - Block scope
  - Package scope
    - Should be small 
  - Outside package
    - Should be capitals

- Casting
  - ```var i float32 = 92.3
    var j int = int(i)```
  - strconv package is used for converting strings to integers.

- Data Types
  - Boolean
    - defaulted to false 
    - var n bool = true
    - var n anotherBool = false
    - n := 1 == 2
  - Numbers
  - Text 
  
