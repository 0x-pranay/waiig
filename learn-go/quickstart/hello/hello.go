package main 

import (
  "fmt"
  "log"
  "example.com/greetings"
)

func main() {
  // log properties
  log.SetPrefix("greeting: ")
  log.SetFlags(0)

  // A slice of names
  names := []string{"Pranay", "Kumar", "Velisoju"}

  // request greeting messages for the names.
  messages, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }

  // message, err := greetings.Hello("Pranay")
  // if err != nil {
  //   log.Fatal(err)
  // }
  fmt.Println(messages)
}
