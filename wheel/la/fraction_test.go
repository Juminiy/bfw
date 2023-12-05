package la

import (
	"bfw/wheel/num"
	"testing"
)

func TestMakeFractionMatrix(t *testing.T) {
	fm := MakeFractionMatrix([][]*num.Fraction{
		{num.MakeFraction(0, 1), num.MakeFraction(1, 2), num.MakeFraction(1, 2)},
		{num.MakeFraction(1, 3), num.MakeFraction(1, 3), num.MakeFraction(1, 3)},
		{num.MakeFraction(0, 1), num.MakeFraction(1, 1), num.MakeFraction(0, 1)}})
	fm.power(fm.makeCopy(), 3)
	fm.Display()
}
