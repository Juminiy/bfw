package char

import (
	"fmt"
	"testing"
)

func TestGetExponent(t *testing.T) {
	fmt.Println(GetSubScript("2"))
	fmt.Println(GetExponent("log2(7)"))
	fmt.Println(GetExponent("log2(3)"))
	fmt.Println(GetExponent("1.6"))
	fmt.Println(GetExponent("k"))
}

func TestGetFraction(t *testing.T) {
	fmt.Println(GetFraction("100", "20"))
}

func TestGetLowerCaseGeekAlphabetBySpell(t *testing.T) {
	fmt.Println(GetLowerCaseGeekAlphabetBySpell("Lambda"))
	fmt.Println(GetUpperCaseGeekAlphabetBySpell("Lambda"))
	fmt.Println(getSuperScript("18") + GetSubScript("18"))
}
