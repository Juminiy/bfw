package char

import (
	"fmt"
	"testing"
)

func TestGetExponent(t *testing.T) {
	fmt.Println(GetExponent("-2345"))
	fmt.Println(GetExponent("-12345"))
}

func TestGetFraction(t *testing.T) {
	fmt.Println(GetFraction("100", "20"))
}

func TestGetLowerCaseGeekAlphabetBySpell(t *testing.T) {
	fmt.Println(GetLowerCaseGeekAlphabetBySpell("Lambda"))
	fmt.Println(GetUpperCaseGeekAlphabetBySpell("Lambda"))
	fmt.Println(getSuperScript("18") + GetSubScript("18"))
}
