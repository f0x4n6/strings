// Package ustrings provides methods for Unicode and ASCII string carving.
//
// Source: https://github.com/robpike/strings/blob/master/strings.go
package ustrings

import (
	"bufio"
	"bytes"
	"strconv"
)

// String data
type String struct {
	// Offset of string
	Offset uint64
	// Value of string
	Value string
}

// Carve Unicode and/or ASCII string data.
//
// The returned channel will be closed at the end of the operation.
func Carve(data []byte, min, max int, ascii bool) <-chan *String {
	ch := make(chan *String, 1024)

	go func() {
		b := bufio.NewReader(bytes.NewReader(data))
		s := make([]rune, 0, max)
		i := uint64(0)

		flush := func() {
			if len(s) >= min {
				v := string(s)
				ch <- &String{i - uint64(len(v)), v}
			}
			s = s[:0]
		}

		defer close(ch)
		defer flush()

		var r rune
		var n int
		var err error

		for ; ; i += uint64(n) {
			if r, n, err = b.ReadRune(); err != nil {
				return
			}

			if !strconv.IsPrint(r) || ascii && r >= 0xFF {
				flush()
				continue
			}

			if len(s) >= max {
				flush()
			}

			s = append(s, r)
		}
	}()

	return ch
}
