package tax

const (
	PersonalTax   uint8 = 1
	EnterpriseTax uint8 = 2
)

type Tax interface {
	Cal(...float64) float64
}

func CalTax(rule [][]float64, taxType uint8, income ...float64) float64 {
	r := MakeRule(rule, makeTaxByType(taxType))
	return r.t.Cal(income...)
}

func makeTaxByType(taxType uint8) Tax {
	switch taxType {
	case PersonalTax:
		return &Personal{}
	case EnterpriseTax:
		return &Enterprise{}
	default:
		panic("unsupported tax type")
	}
}

func inRange(dest, src1, src2 float64) (bool, bool) {
	if src1-src2 > 0 {
		panic("input range error")
	}
	return src1 <= dest && dest <= src2,
		src1 <= dest && src2 <= dest
}

func inRangeV2(dest, src1, src2 float64) bool {
	if src1-src2 > 0 {
		panic("input range error")
	}
	return src1 <= dest && dest <= src2
}
