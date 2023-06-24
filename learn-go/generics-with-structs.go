package main

import "fmt"

type Category struct {
  ID int32
  Name string
  Slug string
}

type Post struct {
  ID int32
  Categories []Category
  Title string
  Description string
  Slug string
}

type cacheable interface {
  Category | Post
}

type cache[T cacheable] struct {
  data map[string]T
}

func (c *cache[T]) Set(key string, value T) {
  c.data[key] = value
}

func (c *cache[T]) Get(key string) (v T) {
  if v, ok := c.data[key]; ok {
    return v
  }
  return
}

func New[T cacheable]() *cache[T] {
  c := cache[T]{}
  c.data = make(map[string]T)
  return &c
}

func main () {
  fmt.Println("program started")

  category := Category{
    ID: 1,
    Name: "Go Generics",
    Slug: "go-generics",
  }

  cc := New[Category]()

  cc.Set(category.Slug, category)

  fmt.Println(*cc)
  fmt.Println(category)

}
