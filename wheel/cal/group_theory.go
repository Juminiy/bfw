package cal

import "bfw/wheel/num"

// FracGroup
// FracGroup is an Example
type FracGroup struct {
	opt rune
	set map[rune]FracElem
}

func (g *FracGroup) E() *FracElem {
	return &FracElem{'e', num.MakeFraction(1, 1)}
}

type FracElem struct {
	char rune
	frac *num.Fraction
}

func (fe *FracElem) Inv() *FracElem {
	return fe.inv()
}

func (fe *FracElem) inv() *FracElem {
	fe.frac = nil
	fe.frac = fe.frac.Inv()
	return fe
}
