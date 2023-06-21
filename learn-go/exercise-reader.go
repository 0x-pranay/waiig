package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// https://go.dev/tour/methods/22

func (r MyReader) Read(b []byte) (int, error) {
  for i := 0; i < len(b); i++ {
    b[i] = byte('A')
  }
  return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

