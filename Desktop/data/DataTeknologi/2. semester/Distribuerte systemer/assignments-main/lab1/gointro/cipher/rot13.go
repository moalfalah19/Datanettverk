package cipher

import (
	"io"
)

// rot13Reader implements io.Reader and applies the ROT13 cipher to the underlying reader's data.
type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	// Read data from the underlying reader into p
	n, err = r.r.Read(p)
	if err != nil && err != io.EOF {
		return n, err
	}

	// Apply ROT13 transformation to the read data
	for i := 0; i < n; i++ {
		switch {
		case p[i] >= 'A' && p[i] <= 'Z': // Uppercase letters
			p[i] = 'A' + (p[i]-'A'+13)%26
		case p[i] >= 'a' && p[i] <= 'z': // Lowercase letters
			p[i] = 'a' + (p[i]-'a'+13)%26
		}
	}

	// Return the number of bytes read and any error (including EOF)
	return n, err
}
