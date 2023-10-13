package lang

import (
	"fmt"
	"testing"
)

func TestQReadInt32(t *testing.T) {
	r := QReadInt32()
	fmt.Println(r)
}
