package num

import (
	"bfw/wheel/adt"
	"errors"
)

var (
	FractionDenominatorCanNotBeZeroError = errors.New("denominator of a fraction cannot be zero")
)

// Fraction
// sign = true, -
// sign = false, +
// IntPair's key = numerator
// IntPair's val = denominator
type Fraction struct {
	sign bool
	adt.IntPair
}

func (frac *Fraction) makeCopy() *Fraction {
	fracCopy := &Fraction{}
	fracCopy.Assign(frac.Self())
	return fracCopy
}

func (frac *Fraction) setSign(sign bool) *Fraction {
	frac.sign = sign
	return frac
}

func (frac *Fraction) setND(n, d int) *Fraction {
	frac.SetKV(n, d)
	return frac
}

func (frac *Fraction) setNOpt(opt rune, lambda int) *Fraction {
	frac.SetKey(int2IntOpt(frac.GetKey(), opt, lambda))
	return frac
}

func (frac *Fraction) setDOpt(opt rune, lambda int) *Fraction {
	frac.SetVal(int2IntOpt(frac.GetVal(), opt, lambda))
	return frac
}

func (frac *Fraction) Sim() *Fraction {
	return frac.makeCopy().sim()
}

func (frac *Fraction) sim() *Fraction {
	gCd := gcd(frac.GetKV())
	frac.setNOpt('/', gCd).
		setDOpt('/', gCd)
	return frac
}

func (frac *Fraction) Add(f *Fraction) *Fraction {
	return frac.makeCopy().add(f)
}

// n1/d1 opt n2/d2 =
// resD = lcm(d1, d2)
// nd1 = resD / d1
// nd2 = resD / d2
// (n1*nd1 opt n2*nd2) / resD
func (frac *Fraction) add(f *Fraction) *Fraction {
	n1, d1 := frac.GetKV()
	n2, d2 := frac.GetKV()
	resD := lcm(d1, d2)
	nd1, nd2 := resD/d1, resD/d2
	return frac.setND(n1*nd1+n2*nd2, resD).sim()
}

func (frac *Fraction) Sub(f *Fraction) *Fraction {
	return frac.makeCopy().sub(f)
}

func (frac *Fraction) sub(f *Fraction) *Fraction {

	return frac
}

func (frac *Fraction) Mul(f *Fraction) *Fraction {
	return frac.makeCopy().mul(f)
}

func (frac *Fraction) mul(f *Fraction) *Fraction {

	return frac
}

func (frac *Fraction) Div(f *Fraction) *Fraction {
	return frac.makeCopy().div(f)
}

func (frac *Fraction) div(f *Fraction) *Fraction {

	return frac
}

func (frac *Fraction) Inv() *Fraction {
	return frac.inv()
}

func (frac *Fraction) inv() *Fraction {
	if frac.GetKey() == 0 {
		panic(FractionDenominatorCanNotBeZeroError)
	}
	frac.SetKVSwap()
	return frac
}
