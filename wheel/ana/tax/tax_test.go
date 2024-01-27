package tax

import (
	"fmt"
	"testing"
)

// 5000, 7500, 6000, 11500, 3500, 6500, 6300, 3700, 6300, 3700
func TestCalTax(t *testing.T) {
	fmt.Println(CalTax([][]float64{
		{0, 0},
		{5000, 0.03},
		{8000, 0.1},
		{17000, 0.2},
		{30000, 0.25},
		{40000, 0.3},
		{60000, 0.35},
		{85000, 0.45},
	}, PersonalTax, 10000))
}
