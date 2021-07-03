// Rot13 rotates each letter by 13 place.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		rot13(os.Stdin)
		return
	}
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		rot13(f)
	}
}

func rot13(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Bytes()
		for i, b := range line {
			if b >= 'A' && b <= 'Z' {
				line[i] = (line[i]-'A'+13)%26 + 'A'
			} else if b >= 'a' && b <= 'z' {
				line[i] = (line[i]-'a'+13)%26 + 'a'
			}
		}
		fmt.Printf("%s\n", line)
	}
}
