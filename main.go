// Carve ASCII and Unicode strings from files.
//
// Usage:
//
//	strings [nmtao] file
//
// The options are:
//
//	n uint
//	    Minimum string length (default 3).
//	m uint
//	    Maximum string length.
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
	"math"
	"os"

	"go.foxforensics.dev/go-mmap"
	"go.foxforensics.dev/strings/strings"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "usage: strings [nmtao] file")
		os.Exit(2)
	}

	x := flag.Uint("n", 3, "minimum string length")
	y := flag.Uint("m", math.MaxUint32, "maximum string length")
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

	for s := range strings.Carve(m, *x, *y, *a, *t) {
		if *o {
			_, _ = fmt.Printf("%08x %s\n", s.Offset, s.Value)
		} else {
			_, _ = fmt.Println(s.Value)
		}
	}
}
