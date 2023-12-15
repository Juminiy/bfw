package dao

import (
	"errors"
	"fmt"
)

const (
	fiveElementFirst = 1 << iota
	fiveElementSecond
	fiveElementThird
	fiveElementFourth
	fiveElementFifth
	zeroRate      float64 = 0.0
	oneQuarter    float64 = 0.25
	halfRate      float64 = 0.5
	threeQuarters float64 = 0.75
	fullRate      float64 = 1.0
)

var (
	pureFireStrList  = []string{"fire"}
	pureWaterStrList = []string{"water"}
	mixedStrList     = []string{"fire", "water"}
)

type element struct {
	elem int
	rate float64
}

func makeZeroElement() *element {
	return &element{}
}

func makeElement(elem int, rate float64) *element {
	e := &element{}
	e.set(elem, rate)
	return e
}

func (e *element) isZero() bool {
	return e.elem == 0 &&
		e.rate == 0
}

func (e *element) set(elem int, rate float64) *element {
	e.elem = elem
	e.rate = rate
	return e
}

func (e *element) setFire(rate float64) *element {
	e.set(fiveElementFirst, rate)
	return e
}

func (e *element) setWater(rate float64) *element {
	e.set(fiveElementSecond, rate)
	return e
}

func (e *element) setPureFire() *element {
	e.setFire(fullRate)
	return e
}

func (e *element) setPureWater() *element {
	e.setWater(fullRate)
	return e
}

func makePureFire() *element {
	return makeElement(fiveElementFirst, fullRate)
}

func makePureWater() *element {
	return makeElement(fiveElementSecond, fullRate)
}

func makePureWood() (*element, *element) {
	return makeElement(fiveElementFirst, threeQuarters),
		makeElement(fiveElementSecond, oneQuarter)
}

func makePureGold() (*element, *element) {
	return makeElement(fiveElementFirst, oneQuarter),
		makeElement(fiveElementSecond, threeQuarters)
}

func makePureEarth() (*element, *element) {
	return makeElement(fiveElementFirst, halfRate),
		makeElement(fiveElementSecond, halfRate)
}

func (e *element) optE(e1, e2 *element, opt rune) {
	switch opt {
	case '+':
		{

		}
	case '-':
		{

		}
	case '*':
		{

		}
	case '/':
		{

		}
	case '^':
		{

		}
	}
}

func (e *element) display(isPrintln bool) *element {
	fmt.Printf("<%d:%f>", e.elem, e.rate)
	if isPrintln {
		fmt.Println()
	}
	return e
}

type Element struct {
	elems map[string]*element
}

func MakeZeroElement() *Element {
	elems := &Element{}
	elems.elems = make(map[string]*element)
	return elems
}

func MakeElement(ks []string, elems []*element) *Element {
	if len(ks) != len(elems) {
		panic(errors.New("keys and elements is not same size"))
	}

	es := MakeZeroElement()
	for idx := 0; idx < len(ks); idx++ {
		es.set(ks[idx], elems[idx])
	}
	return es
}

func (E *Element) isZero() bool {
	for _, elem := range E.elems {
		if !elem.isZero() {
			return false
		}
	}
	return true
}

func (E *Element) set(k string, elem *element) *Element {
	E.elems[k] = nil
	E.elems[k] = elem
	return E
}

func (E *Element) get(k string) *element {
	return E.elems[k]
}

func (E *Element) contribute() *Element {
	return E
}

func (E *Element) distribute() []*element {
	return nil
}

func Fire() *Element {
	fire := makePureFire()
	return MakeElement(pureFireStrList, []*element{fire})
}

func Water() *Element {
	water := makePureWater()
	return MakeElement(pureWaterStrList, []*element{water})
}

func Wood() *Element {
	fire, water := makePureWood()
	return MakeElement(mixedStrList, []*element{fire, water})
}

func Gold() *Element {
	fire, water := makePureGold()
	return MakeElement(mixedStrList, []*element{fire, water})
}

func Earth() *Element {
	fire, water := makePureEarth()
	return MakeElement([]string{"fire", "water"}, []*element{fire, water})
}

func (E *Element) Display() *Element {
	fmt.Printf("[")
	for k, v := range E.elems {
		if !v.isZero() {
			fmt.Printf("(%s:", k)
			v.display(false)
			fmt.Printf(")")
		}
	}
	fmt.Println("]")
	return E
}
