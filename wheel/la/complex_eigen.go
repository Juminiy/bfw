package la

import (
	"bfw/wheel/lang"
	"errors"
	"fmt"
	"log"
)

const (
	eigenValuesNoSize                   int = 0
	eigenValueAlgebraicMultiplicityZero int = 0
	eigenValueAlgebraicMultiplicityOne  int = 1
	eigenValueGeometricMultiplicityZero int = 0
	eigenValueGeometricMultiplicityOne  int = 1
)

var (
	eigenValueNotExists = errors.New("eigen value does not exist")
)

type EigenValues struct {
	multiplicity int
	values       map[complex128]*EigenValue
}

func ConstructEigenValues(slice []complex128) *EigenValues {
	eigenValues := &EigenValues{}
	return eigenValues.Construct(slice)
}

func (evs *EigenValues) Construct(slice []complex128) *EigenValues {
	evs.assign(len(slice))
	for _, lambda := range slice {
		evs.setEigenValue(lambda)
	}
	return evs
}

func (evs *EigenValues) validate() bool {
	if valuesSize := evs.getValuesSize(); valuesSize == eigenValuesNoSize ||
		len(evs.values) != valuesSize {
		return false
	}
	return true
}

func (evs *EigenValues) assign(size int) {
	evs.setValuesValues(make(map[complex128]*EigenValue, size))
	evs.setMultiplicity(size)
}

func (evs *EigenValues) setSelf(evst *EigenValues) {
	evs.setValues(evst.values, evst.multiplicity)
}

func (evs *EigenValues) setValues(values map[complex128]*EigenValue, multiplicity int) {
	evs.setValuesValues(values)
	evs.setMultiplicity(multiplicity)
}

func (evs *EigenValues) setValuesValues(values map[complex128]*EigenValue) {
	evs.values = nil
	evs.values = values
}

func (evs *EigenValues) setMultiplicity(multiplicity int) {
	evs.multiplicity = multiplicity
}

func (evs *EigenValues) getValuesSize() int {
	return len(evs.values)
}

func (evs *EigenValues) get(lambda complex128) *EigenValue {
	return evs.values[lambda]
}

func (evs *EigenValues) set(lambda complex128, ev *EigenValue) {
	evs.values[lambda] = ev
}

func (evs *EigenValues) setEigenValue(lambda complex128) {
	if ev := evs.get(lambda); ev != nil {
		ev.incAlgebraicMultiplicity()
	} else {
		evs.set(lambda, ConstructEigenValue(lambda))
	}
}

func (evs *EigenValues) ValidateMultipleRoots() bool {
	return evs.traverseToGetBool(func(ev *EigenValue) bool {
		return ev.algebraicMultiplicity > 1
	}, true, true, false)
}

func (evs *EigenValues) ValidateAllRealRoots() bool {
	return evs.traverseToGetBool(func(ev *EigenValue) bool {
		return lang.IsComplex128PureReal(ev.lambda)
	}, false, false, true)
}

func (evs *EigenValues) traverseToGetBool(funcPtr func(*EigenValue) bool, predictResult, orResult, andResult bool) bool {
	for _, ev := range evs.values {
		if funcPtr(ev) == predictResult {
			if orResult {
				return predictResult
			}
			if andResult {
				return predictResult
			}
		}
	}
	return !predictResult
}

func (evs *EigenValues) Display(logger ...*log.Logger) {
	valueCnt := 1
	for _, ev := range evs.values {
		ev.Display(valueCnt)
		fmt.Printf(", ")
		valueCnt += ev.algebraicMultiplicity
	}
}

type EigenValue struct {
	lambda                complex128
	algebraicMultiplicity int
	geometricMultiplicity int
	vectors               *EigenVectors
}

func ConstructEigenValue(lambda complex128) *EigenValue {
	eigenValue := &EigenValue{}
	return eigenValue.Construct(lambda)
}

func (ev *EigenValue) Construct(lambda complex128) *EigenValue {
	ev.setLambda(lambda)
	ev.setAM(eigenValueAlgebraicMultiplicityOne)
	ev.setGM(eigenValueGeometricMultiplicityOne)
	return ev
}

func (ev *EigenValue) setSelf(evt *EigenValue) {
	ev.setValues(evt.lambda, evt.algebraicMultiplicity, evt.geometricMultiplicity, evt.vectors)
}

func (ev *EigenValue) setValues(lambda complex128, am, gm int, vectors *EigenVectors) {
	ev.setLambda(lambda)
	ev.setAM(am)
	ev.setGM(gm)
	ev.setVectors(vectors)
}

func (ev *EigenValue) setLambda(lambda complex128) {
	ev.lambda = lambda
}

func (ev *EigenValue) setAM(am int) {
	ev.algebraicMultiplicity = am
}

func (ev *EigenValue) setGM(gm int) {
	ev.geometricMultiplicity = gm
}

func (ev *EigenValue) setVectors(vectors *EigenVectors) {
	ev.vectors = nil
	ev.vectors = vectors
}

func (ev *EigenValue) incAlgebraicMultiplicity() {
	ev.algebraicMultiplicity++
}

func (ev *EigenValue) decAlgebraicMultiplicity() {
	ev.algebraicMultiplicity--
}

// Display
// assume startIndex = 1, am = 1
// display: λ1 = lambda
// assume startIndex = 2, am = 2
// display: λ2 = λ3 = lambda
func (ev *EigenValue) Display(startIndex int) {
	for cnt := startIndex; cnt < startIndex+ev.algebraicMultiplicity; cnt++ {
		fmt.Printf(string(EigenPolyMatrixDefaultAES)+"%d = ", cnt)
	}
	lang.DisplayComplex128(5, 5, ev.lambda)
}

type EigenVectors struct {
	VectorGroup
	geometricMultiplicity int
}
