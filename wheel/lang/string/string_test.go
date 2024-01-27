package string

import (
	"fmt"
	"testing"
)

// go test -bench="$Dup$" -cpu=1,2,4,10  -benchmem .

func BenchmarkStringBuilderPreAlloc(b *testing.B) {
	Dup(b, dupStringV7)
}

func BenchmarkByteSlicePreAlloc(b *testing.B) {
	Dup(b, dupStringV6)
}

func BenchmarkByteBufferPreAlloc(b *testing.B) {
	Dup(b, dupStringV8)
}

func BenchmarkStringBuffer(b *testing.B) {
	Dup(b, dupStringV4)
}

func BenchmarkStringBuilder(b *testing.B) {
	Dup(b, dupStringV3)
}

func BenchmarkByteSlice(b *testing.B) {
	Dup(b, dupStringV5)
}

func BenchmarkRawString(b *testing.B) {
	Dup(b, dupStringV1)
}

func BenchmarkFMTString(b *testing.B) {
	Dup(b, dupStringV2)
}

func TestDup(t *testing.T) {
	fmt.Printf("%06b\n", 0b101010^0)
	fmt.Printf("%06b\n", 0b010101^0)
}
