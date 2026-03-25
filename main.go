// ASCII and Unicode string carving tool.
//
// Usage:
//
//	strings file
//
// The arguments are:
//
//	file
//	    File to carve (required).
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func stream(name string, ch chan<- byte) {
	f, err := os.Open(name)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	defer func() { _ = f.Close() }()

	buf, err := io.ReadAll(f)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, b := range buf {
		ch <- b
	}

	close(ch)
}

func flush(runes []rune) []rune {
	if len(runes) >= 3 {
		s := string(runes)

		if len(strings.TrimSpace(s)) > 0 {
			_, _ = fmt.Println(s)
		}
	}

	return runes[:0]
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, "usage: strings FILE")
		os.Exit(2)
	}

	ch := make(chan byte, 1024)

	go stream(os.Args[1], ch)

	buf := make([]byte, 4)

	var runes []rune

	for b := range ch {
		buf[0] = b

		l := 1
		k := 1

		if b&0x80 == 0 {
			k = 1
		} else if b&0xE0 == 0xC0 {
			k = 2
		} else if b&0xF0 == 0xE0 {
			k = 3
		} else if b&0xF8 == 0xF0 {
			k = 4
		}

		if k > 1 {
			for i := 1; i < k; i++ {
				if b, ok := <-ch; ok {
					buf[i] = b
				} else {
					break
				}

				l++
			}
		}

		r, _ := utf8.DecodeRune(buf[:l])

		if r != utf8.RuneError && unicode.IsPrint(r) {
			runes = append(runes, r)
		} else {
			runes = flush(runes)
		}
	}

	flush(runes)
}
