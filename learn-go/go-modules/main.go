package main

import (
  "fmt"
  "hello/greeting"
)

func main() {
  fmt.Print("calling from main \n");
  greeting.greet();
}
