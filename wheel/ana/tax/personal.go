package tax

import (
	"fmt"
)

type Personal struct {
	inc float64
	tax float64
	R   *Rule `tax_cal:"rule"`
}

// Cal
// Cal tax by Rule
func (t *Personal) Cal(income ...float64) float64 {
	return t.set(t.R.calV2(income...), income...)
}

func (t *Personal) set(tax float64, income ...float64) float64 {
	t.inc = 0
	for i := 0; i < len(income); i++ {
		t.inc += income[i]
	}
	t.tax = tax
	return t.tax
}

func (t *Personal) Print() {
	fmt.Printf("income = %.2f, tax = %.2f", t.inc, t.tax)
}
