// Package ustrings provides methods for Unicode and ASCII string carving.
//
// Source: https://github.com/robpike/strings/blob/master/strings.go
package ustrings

import (
	"bufio"
	"bytes"
	"strconv"
)

// Carve Unicode and/or ASCII strings.
func Carve(data []byte, min, max int, ascii bool, flush func(int64, []rune)) {
	b := bufio.NewReader(bytes.NewReader(data))
	s := make([]rune, 0, max)
	i := int64(0)

	reset := func() {
		if len(s) >= min {
			flush(i-int64(len(string(s))), s)
		}
		s = s[:0]
	}

	defer reset()

	var r rune
	var n int
	var err error

	for ; ; i += int64(n) {
		if r, n, err = b.ReadRune(); err != nil {
			return
		}

		if !strconv.IsPrint(r) || ascii && r >= 0xFF {
			reset()
			continue
		}

		if len(s) >= max {
			reset()
		}

		s = append(s, r)
	}
}
