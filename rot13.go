// Rot13 rotates each letter by 13 place.
package main

import (
	"io"
	"os"
)

func main() {
	r := Rot13{os.Stdin}
	io.Copy(os.Stdout, r)
}

type Rot13 struct {
	r io.Reader
}

func (r Rot13) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	for i := 0; i < n; i++ {
		if 'A' <= buf[i] && buf[i] <= 'Z' {
			buf[i] = 'A' + (buf[i]-'A'+13)%26
		} else if 'a' <= buf[i] && buf[i] <= 'z' {
			buf[i] = 'a' + (buf[i]-'a'+13)%26
		}
	}
	return n, err
}
