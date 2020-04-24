package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b < 'A'+13:
		return b + 13
	case 'A'+13 <= b && b < 'A'+26:
		return b - 13
	case 'a' <= b && b < 'a'+13:
		return b + 13
	case 'a'+13 <= b && b < 'a'+26:
		return b - 13
	default:
		return b
	}
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
