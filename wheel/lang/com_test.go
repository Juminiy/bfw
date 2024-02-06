package lang

import (
	"fmt"
	"testing"
)

func TestComplex128Slice_Len(t *testing.T) {
	a, b := 1+2i, 5+6i
	fmt.Println(Eq(a/b, Div(a, b)))
}

func BenchmarkAbsInt(b *testing.B) {
	n := 1 << 16
	l, r := GenComplex128Slice(n)
	res := make([]complex128, n)
	b.Run("Golang runtime:", func(b *testing.B) {
		for i := range l {
			res[i] = l[i] / r[i]
		}
	})

	b.Run("my Div:", func(b *testing.B) {
		for i := range l {
			res[i] = l[i] / r[i]
		}
	})
}
