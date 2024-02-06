package lang

import (
	"fmt"
	"testing"
)

func TestDisplayComplex128(t *testing.T) {
	f := func() {}
	fmt.Printf("instack: %p\n", f)
	dup(f)
}
