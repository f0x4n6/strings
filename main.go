// Carve Unicode and ASCII strings from binary files.
//
// Usage:
//
//	ustrings [nmao] file
//
// The options are:
//
//	n int
//	    Minimum string length (default 4).
//	m int
//	    Maximum string length (default 255).
//	a
//	    Only ASCII strings.
//	o
//	    Show file offset.
//
// The arguments are:
//
//	file
//	    File to carve (required).
package main

import (
	"flag"
	"fmt"
	"os"

	"go.foxforensics.dev/go-mmap"
	"go.foxforensics.dev/ustrings/ustrings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, "usage: ustrings [nmao] file")
		os.Exit(2)
	}

	x := flag.Int("n", 4, "minimum string length")
	y := flag.Int("m", 255, "maximum string length")
	a := flag.Bool("a", false, "only ASCII strings")
	o := flag.Bool("o", false, "show file offset")

	flag.Parse()

	f, err := os.Open(flag.Arg(0))

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() { _ = f.Close() }()

	m, err := mmap.Map(f, mmap.RDONLY, 0)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() { _ = m.Unmap() }()

	ustrings.Carve(m, *x, *y, *a, func(i int64, r []rune) {
		if *o {
			_, _ = fmt.Printf("%08x %s\n", i, string(r))
		} else {
			_, _ = fmt.Println(string(r))
		}
	})
}
