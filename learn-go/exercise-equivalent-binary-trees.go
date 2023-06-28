package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func Walk(t *Tree, ch chan int) {
	// In order traversal.
	if t != nil {
		Walk(t.Left, ch)
		// fmt.Println(t.Value)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Same(t1 *Tree, t2 *Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func main() {
	tree1 := New(1)
	ch := make(chan int, 10)
	fmt.Println(tree1)
	Walk(tree1, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	t1 := New(1)
	t2 := New(1)
	r := Same(t1, t2)
	fmt.Println(r)

	t3 := New(1)
	t4 := New(2)
	r = Same(t3, t4)
	fmt.Println(r)
}
