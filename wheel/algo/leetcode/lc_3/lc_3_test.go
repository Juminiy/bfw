package lc_3

import (
	"fmt"
	"testing"
)

func TestLC3(t *testing.T) {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi(" -42"))
	fmt.Println(myAtoi("   4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("00000-42a1234"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("9223372036854775808"))
}
