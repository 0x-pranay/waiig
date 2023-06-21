package main

import (
  "fmt"
  "io"
  "strings"
  "bytes"
)

func main() {

  usingStrings()

  usingBuffers()
}

func usingStrings() {
  r := strings.NewReader("Hello, Reader1234")

  b := make([]byte, 8)

  for {
    n, err := r.Read(b)
    fmt.Printf("n = %v, err = %v, b = %v \n", n, err, b)

    fmt.Printf("b[:n] = %s\n", b[:n])
    if err == io.EOF {
      break
    }
  }
}

func usingBuffers() {
  fmt.Println("using Buffers....")

  r := bytes.NewReader([]byte("Hello, Reader"))

  fmt.Printf("bufferLen: %v", r.Len())

  rInt := bytes.NewReader([]byte{11, 12,13, 255, 255})

  b := make([]byte, 10)
  
  for {
    n, err := rInt.Read(b)
    fmt.Printf("n = %v, err = %v, b = %v \n", n, err, b)
    fmt.Printf("b[:n] = %d\n", b[:n])
    if err == io.EOF {
      break
    }
  }
  fmt.Printf("bufferLen: %v", rInt.Len())

}
