// Carve Unicode and/or ASCII strings from files.
//
// Usage:
//
//	ustrings [nmtao] file
//
// The options are:
//
//	n uint
//	    Minimum string length (default 4).
//	m uint
//	    Maximum string length (default 256).
//	t
//		Trim spaces from both ends.
//	a
//	    Only ASCII strings.
//	o
//	    Show file offset.
//
// The arguments are:
//
//	file
//	    File to be carved (required).
package main

import (
	"flag"
	"fmt"
	"os"

	"go.foxforensics.dev/go-mmap"
	"go.foxforensics.dev/ustrings/ustrings"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "usage: ustrings [nmtao] file")
		os.Exit(2)
	}

	x := flag.Uint("n", 4, "minimum string length")
	y := flag.Uint("m", 256, "maximum string length")
	t := flag.Bool("t", false, "trim spaces from both ends")
	a := flag.Bool("a", false, "only ASCII strings")
	o := flag.Bool("o", false, "show file offset")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

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

	for s := range ustrings.Carve(m, *x, *y, *t, *a) {
		if *o {
			_, _ = fmt.Printf("%08x %s\n", s.Offset, s.Value)
		} else {
			_, _ = fmt.Println(s.Value)
		}
	}
}
