package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1,2})

  z := Vertex{ 2, 4 }

  fmt.Println(z.X)

  z.X = 10

  fmt.Println(z.X)
  
  fmt.Println("Pointer to a struct")
  
  p := &z
  p.X = 1e9
  fmt.Println(z)

  fmt.Println("Struct Literal")

  v1 := Vertex{ X: 1, Y: 10 }
  fmt.Println(v1)

  p1 := &v1
  fmt.Println(p1)
}
