package ustrings

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

var bin = filepath.Join("..", "testdata", "test.bin")
var txt = filepath.Join("..", "testdata", "test.txt")

func TestCarve(t *testing.T) {
	for _, tt := range []struct {
		name  string
		path  string
		ascii bool
		count int
	}{
		{
			name:  "ASCII",
			path:  bin,
			ascii: true,
			count: 3117,
		},
		{
			name:  "Unicode",
			path:  txt,
			ascii: false,
			count: 582,
		},
	} {
		t.Run("Test Carve "+tt.name, func(t *testing.T) {
			buf, err := fixture(tt.path)

			if err != nil {
				t.Fatalf("Carve: %v", err)
			}

			n := 0

			for range Carve(buf, 4, 255, tt.ascii) {
				n++
			}

			if n != tt.count {
				t.Fatal("count mismatch")
			}
		})
	}
}

func BenchmarkCarve(b *testing.B) {
	b.Run("Benchmark Carve", func(b *testing.B) {
		bin, err := fixture(bin)

		if err != nil {
			b.Fatalf("Carve: %v", err)
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			for range Carve(bin, 4, 255, false) {
			}
		}
	})
}

func fixture(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	b, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return b, nil
}
