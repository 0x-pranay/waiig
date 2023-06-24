package main

// List represents a singly-linked list that holds
// values of any type.
import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) append(c *List[T]) *List[T] {
	l.next = c
	return l
}

// TODO: display method


func emptyNode[T any]() *List[T]{
	n := List[T]{ next: nil}
	return &n
}

func newNode[T any](v T) *List[T] {
	n := List[T]{ next: nil, val: v}
	return &n
}



func main() {
	emptyIntNode := emptyNode[int]()
	fmt.Println(emptyIntNode)
	
	rootStringLL := newNode[string]("Hello")
	fmt.Println(*rootStringLL)
	
	child1 := newNode[string](" World!")
	
	rootStringLL.append(child1)
	fmt.Println(*rootStringLL)
	
	
}

