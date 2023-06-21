package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {

	n, err := rot13.r.Read(b)
	if err == io.EOF {
		return 0, io.EOF
	}
	// fmt.Printf("n: %v, b1: %q \n", n, b1)

	for i := 0; i < n; i++ {
		// b[i] = rot13transform(b1[i])
		char := b[i]
		switch {
		case char >= 'A' && char <= 'Z':
			b[i] = 'A' + ((char-'A')+13)%26
		case char >= 'a' && char <= 'z':
			b[i] = 'a' + ((char-('a'))+13)%26
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

