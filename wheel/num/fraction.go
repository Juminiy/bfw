package num

import (
	"bfw/wheel/adt"
	"bfw/wheel/lang"
	"errors"
	"fmt"
)

var (
	FractionDenominatorCanNotBeZeroError = errors.New("denominator of a fraction cannot be zero")
)

// Fraction
// sign = true, -
// sign = false, +
// IntPair's key = numerator -> n; n >= 0
// IntPair's val = denominator -> d; d > 0
// n,d is always positive
// sign is field sign
type Fraction struct {
	sign bool
	adt.IntPair
}

func MakeFraction(n, d int) *Fraction {
	frac := &Fraction{}
	frac.setND(n, d).format().
		sim()
	return frac
}

func (frac *Fraction) validate() bool {
	n, d := frac.GetKV()
	return d > 0 && n >= 0
}

// format
// only used for set n,d and reset n,d
func (frac *Fraction) format() *Fraction {
	if n, d := frac.GetKV(); d != 0 {
		if !lang.IsIntSameSign(n, d) {
			frac.setSign(true)
		} else {
			frac.setSign(false)
		}
		if n == 0 {
			frac.setSign(false)
		}
		frac.setND(lang.AbsInt(n), lang.AbsInt(d))
	} else {
		panic(FractionDenominatorCanNotBeZeroError)
	}
	return frac
}

func (frac *Fraction) makeCopy() *Fraction {
	fracCopy := &Fraction{}
	fracCopy.setSelf(frac.GetKey(), frac.GetVal(), frac.sign)
	return fracCopy
}

func (frac *Fraction) setSelf(n, d int, sign ...bool) {
	frac.setSign(sign...)
	frac.setND(n, d)
}

func (frac *Fraction) setSign(sign ...bool) *Fraction {
	if signLen := len(sign); signLen > 0 {
		switch signLen {
		case 1:
			{
				frac.sign = sign[0]
			}
		case 2:
			{
				if sign[0] == sign[1] {
					frac.sign = false
				} else {
					frac.sign = true
				}
			}
		}
	} else {
		frac.sign = false
	}
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

// n1/d1 opt n2/d2 =
// resD = lcm(d1, d2)
// nd1 = resD / d1
// nd2 = resD / d2
// (n1*nd1 opt n2*nd2) / resD
func (frac *Fraction) opt(f *Fraction, opt rune) *Fraction {
	n1, d1 := frac.GetKV()
	n2, d2 := f.GetKV()
	n1, n2 = lang.GetOriginNum(n1, frac.sign), lang.GetOriginNum(n2, f.sign)
	resD := lcm(d1, d2)
	nd1, nd2 := resD/d1, resD/d2
	switch opt {
	case '+':
		{
			frac.setND(n1*nd1+n2*nd2, resD).format().sim()
		}
	case '-':
		{
			frac.setND(n1*nd1-n2*nd2, resD).format().sim()
		}
	default:
		{
			panic(errors.New("unsupported operator"))
		}
	}
	return frac
}

func (frac *Fraction) Add(f *Fraction) *Fraction {
	return frac.makeCopy().add(f)
}

func (frac *Fraction) add(f *Fraction) *Fraction {
	return frac.opt(f, '+')
}

func (frac *Fraction) Sub(f *Fraction) *Fraction {
	return frac.makeCopy().sub(f)
}

func (frac *Fraction) sub(f *Fraction) *Fraction {
	return frac.opt(f, '-')
}

func (frac *Fraction) Mul(f *Fraction) *Fraction {
	return frac.makeCopy().mul(f)
}

func (frac *Fraction) mul(f *Fraction) *Fraction {
	n1, d1 := frac.GetKV()
	n2, d2 := f.GetKV()
	return frac.setND(n1*n2, d1*d2).sim().setSign(frac.sign, f.sign)
}

func (frac *Fraction) Div(f *Fraction) *Fraction {
	return frac.makeCopy().div(f)
}

func (frac *Fraction) div(f *Fraction) *Fraction {
	return frac.mul(f.inv())
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

func (frac *Fraction) Float64() float64 {
	return float64(frac.GetKey()) / float64(frac.GetVal())
}

func (frac *Fraction) Display(isPrintln ...bool) *Fraction {
	if !frac.validate() {
		panic(FractionDenominatorCanNotBeZeroError)
	}
	if frac.sign {
		fmt.Printf("-")
	}
	fmt.Printf("%d/%d", frac.GetKey(), frac.GetVal())
	if len(isPrintln) > 0 && !isPrintln[0] {

	} else {
		fmt.Println()
	}
	return frac
}
