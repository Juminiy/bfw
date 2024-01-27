package tax

type Enterprise struct {
	TaxRule *Rule `tax_cal:"rule"`
}

func (t *Enterprise) Cal(...float64) float64 {
	return 0.0
}
