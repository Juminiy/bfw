package tax

import (
	"bfw/wheel/lang"
	"math"
	"reflect"
	"sort"
)

var (
	outOfBoundIncomeS = math.Inf(-1)
	outOfBoundIncomeE = math.Inf(1)
)

type (
	Rule struct {
		tax map[float64]*ruleInfo
		inc []float64
		t   Tax
	}
	ruleInfo struct {
		index     int
		ins, ine  float64
		rate, tot float64
	}
)

func MakeRule(oR [][]float64, t Tax) *Rule {
	rule := &Rule{}
	rule.set(oR)
	rule.t = t

	// runtime set value
	tType := reflect.TypeOf(t).Elem()
	tValue := reflect.ValueOf(t).Elem()
	switch tType.Kind() {
	case reflect.Struct:
		{
			for i := 0; i < tType.NumField(); i++ {
				if tType.Field(i).Tag.Get("tax_cal") == "rule" {
					tValue.Field(i).Set(reflect.ValueOf(rule))
				}
			}
		}
	default:
		panic("unsupported tax type")
	}

	//switch t.(type) {
	//case *Personal:
	//	t.(*Personal).r = rule
	//case *Enterprise:
	//	t.(*Enterprise).taxRule = rule
	//default:
	//	panic("unsupported tax type")
	//}
	return rule
}

// originalRule n*2
// [i][0] inc[i]
// [i][1] tax[inc[i]]
// i taxI[tax[inc[i]]]
func (r *Rule) set(originalRule [][]float64) {
	r2da := lang.Real2DArray(originalRule)
	sort.Sort(r2da)
	r.inc = lang.GetReal2DArrayColumn(r2da, 0)
	rl := len(r.inc)
	r.tax = make(map[float64]*ruleInfo, rl+2)
	rate := lang.GetReal2DArrayColumn(r2da, 1)
	tot := 0.0
	r.tax[0] = &ruleInfo{index: 0, ins: outOfBoundIncomeS, ine: 0, rate: 0, tot: 0}
	for i := 0; i < rl; i++ {
		ins, ine, rt := r.inc[i], outOfBoundIncomeE, rate[i]
		if i == rl-1 {
			ine = outOfBoundIncomeE
		} else {
			ine = r.inc[i+1]
		}
		tot += (ine - ins) * rt
		r.tax[r.inc[i]] = &ruleInfo{index: i, ins: ins, ine: ine, rate: rt, tot: tot}
	}
	r.tax[outOfBoundIncomeE] = &ruleInfo{index: rl, ins: outOfBoundIncomeE, ine: outOfBoundIncomeE, rate: 0, tot: 0}
}

func (r *Rule) key(dest float64) float64 {
	for k, v := range r.tax {
		if inRangeV2(dest, v.ins, v.ine) {
			return k
		}
	}
	return outOfBoundIncomeE
}

func (r *Rule) get(key float64) *ruleInfo {
	if key < 0 {
		return r.tax[0]
	}
	return r.tax[key]
}

func (r *Rule) cal(dest float64) float64 {
	key := r.key(dest)
	val := r.get(key)
	preKey := r.key(key - 1)
	preVal := r.get(preKey)
	return preVal.tot + (dest-key)*val.rate
}

func (r *Rule) calQ2(dest ...float64) []float64 {
	tax := make([]float64, len(dest))
	for i := 0; i < len(dest); i++ {
		tax[i] = r.cal(dest[i])
	}
	return tax
}

func (r *Rule) calV2(dest ...float64) float64 {
	tax := r.calQ2(dest...)
	tot := 0.0
	for i := 0; i < len(tax); i++ {
		tot += tax[i]
	}
	return tot
}

// Personal Rule
// (-inf,     0],  0.00
// (0, 	   5000],  0.00
// (5000,  8000],  0.03
// (8000,  17000], 0.10
// (17000, 30000], 0.20
// (30000, 40000], 0.25
// (40000, 60000], 0.30
// (60000, 85000], 0.35
// (85000,  +inf), 0.45

// Enterprise Rule
