# Learning Go

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

- Similar to js objects.
- The zero value of a map is `nil`
- use `make` function to create a map of given type.
- map keys can be any datatype which can be comparate

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
- struct fields are accessed using dot.
```
type Token struct {
  Type string
  Literal string
}

func main() {
  fmt.Println(Token{ "IDENT", "5" })

  tok := Token{ "VARIABLE_DECL", "let" }
  tok.Literal = "const"
  fmt.Println(tok.Literal)
}
```

### Functions
- Function in go are first class functions. meaning functions are values too. They can be assigned to a variable, passed as an argument or a function can return another function. 
- Go functions may be closures. 
- Closure is a function value that references variables from outside its body.

### Methods
- GO does not have classes. However, we can define methods on Types. 
- A method is a function with special receiver argument. 
- We can declare methods on non-struct types, too.

#### Pointer receivers
- Methods with pointer receiver can modify the value of which the reciever points.
These are more common than value recievers. 
```
type Vertex struct {
  X, Y float64
}

// value receiver
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*V.X + v.Y * v.Y)
}

// Pointer reciever
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

```

- Methods with pointer recievers take either a value or a pointer as the receiver.
```
var v Vertex
v.Scale(5) // OK

p := &v
p.Scale(10) // OK
```

- The same happens for value receiver methods, they can take either apointer or a value as the receiver. 

- Two reasons to choose pointer reciever methods
  1. So that method can modify its receivers value.
  2. To avoid copying the value on each method call. This can be more efficient if the receiver is a large struct


### Interfaces
- An interface type is defined as a set of method signatures.
- The value of interface type can hold any value that implements those methods. 
- Interfaces are implemented implicitly
  - We dont need to use `implements` keyword as in other languages.

```
type Shape interface {
  Area() float64
}

type Rect struct {
  Length float64
  Breadth float64
}

func (t T) Area() {
  return t.Length * t.Breadth
}

func main() {
  var rectangle Shape = Rect{ 1.0, 2.0 }
  rectange.Area()

}
```
- Under the hood, interface values can be thought of as a tuple of a value and a concrete type. 
    `(value, type)`

- calling a method on a interface value executes the method of the same name on its underlying type. https://go.dev/tour/methods/11

- If the concreate value inside the interface is nil, th method will be called with a `nil` receiver. 
```
// gracefully handle being called with a nil receiver

func (t *T) M() {
  if t == nil {
      fmt.Println("<nil>")
      return
    }
    fmt.Println(t.S)
}
```
- A nil interface value holds neither value nor concrete type. So calling a method on a nil interface is a run time error because there is no type inside the interface to indicate which concrete method to call. 

#### Empty interface
- interface type with zero methods
- an empty interface may hold values of any type. (Every type implements at least zero methods.)
- Empty interfaces are used by code that handles values of unknown type.

#### Type Assertion
- A **type assertion** provides access to an interface value's underlying concrete value.
```
t := i.(T)
```
- This statements asserts that the interface value **i** holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
- If `i` does not hold a `T`, the statement will trigger a panic. 
- A type assertion can return two values. --> To test whether an interface value holds a specific type. 
```
t, ok := i.(T)
```
- If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true
- Else `t` = zero value of type `T` and `ok` = false and no panic. 
```
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

```

#### Type switches
- A type switch is a construct that permits several type assertions in series.


- A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

Ex: 
```
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```
- https://go.dev/tour/methods/16

#### Stringers

```
type Stringer interface {
  String() string
}
```
- A `Stringer` is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

- Stringer is implemented by any value that has a String method, which defines the “native” format for that value. The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print.
```
import (
	"fmt"
)

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
```

#### Errors

The `error` type is a built-in interface similar to `fmt.Stringer`
```
type error interface {
  Error() string
}
```
- Go program express error state using `error` value. (As with fmt.Stringer, the fmt package looks for the error interface when printing error values. )

- if err != nil  --> denotes success.

- https://go.dev/tour/methods/20 

#### Readers
The `io` package specifies the `io.Reader` interface which represents the read end of a stream of data.

Other implementation of this interface, includes files, network connections, compressors, ciphers and others. 

- `io.Reader` interface has `Read` method.

```

func (T) Read(b []byte) (n int, err error) 

```
`Read` populates the given type slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.


Example - [reader.go](./reader.go)

- https://go.dev/tour/methods/21


### Images [TODO](https://go.dev/tour/methods/24)

## Generics

### Type parameters

Go functions can be writtten to work with multiple types using type parameters. 

This type parameter appears between brackets, before the function arguments. 

Syntax:
```
func Index[T comparable] (s []T, x T) int
```

This declaration means
  `s`   is a slice of type `T` that fulfills built-in constraint `comparable`
  `x`   is also a value of type `T`

  `comparable`  this is a built-in useful constraint that makes using `==` and `!=` operators possible.


Example:
```
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
```

These above function which supports generic types are called **Generic Function**

Generic functions = Declates both ordinary function parameters + **type parameters**

### Declare type constraint as an Interface

```
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```
You declare a type constraint as an interface. The constraint allows any type implementing the interface. For example, if you declare a type constraint interface with three methods, then use it with a type parameter in a generic function, type arguments used to call the function must have all of those methods.

```
type Number interface {
    int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

## Concurrency

### Goroutines

- A lightweight thread managed by the Go runtime
  `go f(x, y, z)`
- evaluation of f,x,y and z heppens in the current thread and the execution of f happends in the new goroutine.

####  Channels

A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is nil.

The optional <- operator specifies the channel direction, send or receive.

```
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

A new, initialized channel value can be made using the built-in function make, which takes the channel type and an optional capacity as arguments:
`make(chan int, 100)`

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A nil channel is never ready for communication.

A channel may be closed with the built-in function `close`.

```
c := make(chan int, 10)
c <- 5
// do something

close(c) // closed the channel only sender can close the channel.

c <-10 // will cause panic, cause channel is already closed.
```

##### Send

```
ch <- 3  // send value 3 to channel ch
```
- A send on an unbuffered channel can proceed if a receiver is ready. 
- A send on a buffered channel can proceed if there is room in the buffer. 
- A send on a closed channel proceeds by causing a run-time panic. 
- A send on a nil channel blocks forever.

##### Receive
The expression blocks until a value is available. Receiving from a nil channel blocks forever. A receive operation on a closed channel can always proceed immediately, yielding the element type's zero value after any previously sent values have been received.

```
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // wait until clock pulse and discard received value
```

With additional untyped boolean  result reporting whether the communication succeeded.
```
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```
The value of ok is
  true if the value received was delivered by a successful send operation to the channel,
  or
  false if it is a zero value generated because the channel is closed and empty.


### Select

- this statements lets a goroutine wait on multiple communication operations. 
- A `select` blocks until one of its cases can run, then it executes that case.
- It chooses one at random if multiple are ready. ( TODO: further reading required)
- The `default` case in a select is run if no other case is ready.

- Use a `default` case to try a send or receive without blocking

---

# Standard library

## fmt
- https://pkg.go.dev/fmt#Stringer
- https://pkg.go.dev/fmt@go1.20.5
### verbs


#### General

- `%v`    the value in default format.When printing structs, the plus flag (%+v) adds field names
- `%#v`  a Go-syntax representation of the value
- `%T`   Type of the value
- `%%`   Used to escape percent sign. a literal percent sign.

#### Boolean
- `%t`    Boolean. The word true or false 

#### Integer



#### String or Slice of bytes

`%s`    the uninterpreted bytes of string or slice
`%q`    a double-quoted string safely escaped with Go syntax

## Individual topics

- Using generics with structs - https://itnext.io/how-to-use-golang-generics-with-structs-8cabc9353d75
- Go concurrency patterns  by Rob Pike, Google I/O 2012 - https://www.youtube.com/watch?v=f6kdp27TYZs



### Go byte
- https://zetcode.com/golang/byte/
