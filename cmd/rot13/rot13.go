// Rot13 rotates each letter by 13 place.
package main

import (
	"io"
	"os"

	"github.com/bbriano/rot13"
)

func main() {
	r := rot13.Rot13{os.Stdin}
	io.Copy(os.Stdout, r)
}
