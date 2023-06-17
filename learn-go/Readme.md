Learning Go

- For primitive datatypes we pass a variable by value instead of its reference in function calls. 

- Go's basic types are
  - bool
  - string
  - int  int8  int16  int32  int64
  - uint uint8 uint16 uint32 uint64 uintptr
  - byte // alias for uint8
  - rune // alias for int32
     // represents a Unicode code point
  - float32 float64
  - complex64 complex128

- variable declarations may be "factored" into blocks, as with import statements.
- The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type. 

- Variables declared without an explicit initial value are given their zero value.
  - 0 for numerica types
  - false for the boolean type and 
  - ""(the empty string) for strings

- Type conversion. The expression T(v) converts the value of v to the type T.
  ex:
  ```
  i := 42
  f := float64(i)
  u := uint(f)
  ```

### Control Statements

### For
- Go has only one looping constuct, the `for` loop
```
for INIT_STATEMT; CONDITION_EXPRESSION; POST_STATEMENT {

}

sum := 0
for i := 0; i< 10; i++ {
  sum += i
}

```
- Init and Post statements are optional. We can drop the semicolons and only have one conditional_expression. This is similar to while loop.
- we can have a forever running loop. 
```
for {

}
```

### If
```
if x < 0 {
  // do something
}
```
- Like in `for`, we can execute a short statement before the condition
- variables declared in these statements are only in scope until the end of the `if` and also available in `else` block

### Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

```
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

### Defer
- A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
- Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.


### Array
- Arrays are fixed size groups of variables of same type
```
var myInts [10]int

primes := [6]int{2, 3, 5, 7, 11, 13}

words := [3]string{
  "Hello",
  "World",
  "Hello world",
  }

```


### Slices
- Slices reference array under the hood. 
- Slices are dynamic arrays.
- create a new slice using `make` keyword.
  ```
  // func make([]T, len, cap) []T
  mySlice := make([]int, 5, 10)

  // The capacity argument is ususally omitted and defaults to the length
  mySlice := make([]int, 5)
  ```
- Slices creatd with `make` will be filled with the zero valued of the type
- If we want to create a slice with a specific set of values, we can use a slice literal:
  ```
  mySlice := []string{"Here", "three", "words"}
  fmt.Println("length of slice", len(mySlice))
  fmt.Println("Capacity of slice", cap(mySlice))
  ```

- `append` function is a variadic function. `func append(slice []Type, elems ...Type) []Type`
- `a[inclusive_low : exclusive_high]`

### Variadic functions and Spread operator

- Variadic functions are function which can take any number of arguments of the same type
- Spread operator is used to multiple arguments from a single array.

```
func printStrings(strings ...string){
  for i := 0; i < len(strings); i++ {
    fmt.Println(strings[i])
  }
}

func main() {
  names := []string{"bob", "sue", "alice"}
  printStrings(names...)
}

```

### Range
- a syntactic sugar to iterate easily over elements of a slice.
```
for INDEX, ELEMENT := range SLICE {

}
```

### Map
- Similar to js objects. The zero value of a map is `nil`
```
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 22


// or

ages := mapp[string]int{
  "John": 37,
  "Mary": 22,
}
```
- `len()` function returns number of key-value pairs.

- Mutation operations: insert, get, delete, check if a key exists. 

### Pointers
- A pointer holds memory address of a value. 
- `*T` is a pointer to a T value. Its zero value is `nil`.
```
var p *int

i := 42

p_i = &i

fmt.Println(*p_i) // reads i through the pointer p_i
*p = 21 // sets i through the pointer p. AKA "dereferencing" or "indirecting"
```
- Unlike C, Go has no pointer arithmetic.

### Structs
- A `struct` is a collection of fields



