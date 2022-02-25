// Rot13 rotates each letter by 13 place.
package rot13

import (
	"io"
)

type Rot13 struct {
	io.Reader
}

func (r Rot13) Read(buf []byte) (int, error) {
	n, err := r.Reader.Read(buf)
	for i := 0; i < n; i++ {
		if 'A' <= buf[i] && buf[i] <= 'Z' {
			buf[i] = 'A' + (buf[i]-'A'+13)%26
		} else if 'a' <= buf[i] && buf[i] <= 'z' {
			buf[i] = 'a' + (buf[i]-'a'+13)%26
		}
	}
	return n, err
}
