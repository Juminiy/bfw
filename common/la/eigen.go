package la

import (
	"errors"
)

const (
	eigenValuesNoSize int = 0
)

var (
	eigenValueNotExists = errors.New("eigen value does not exist")
)

type EigenValues struct {
	size         int
	multiplicity int
	values       map[complex128]*EigenValue
}

func ConstructEigenValues(slice []complex128) *EigenValues {
	eigenValues := &EigenValues{}
	return eigenValues.Construct(slice)
}

func (evs *EigenValues) Construct(slice []complex128) *EigenValues {
	for _, lambda := range slice {
		evs.setEigenValue(lambda)
	}
	return evs
}

func (evs *EigenValues) validate() bool {
	if evs.size == eigenValuesNoSize ||
		len(evs.values) != evs.size {
		return false
	}
	return true
}

func (evs *EigenValues) assign(size int) {
	evs.setSize(size)
	evs.setValuesValues(make(map[complex128]*EigenValue, size))
}

func (evs *EigenValues) setSelf(evst *EigenValues) {
	evs.setValues(evst.values, evst.size)
}

func (evs *EigenValues) setValues(values map[complex128]*EigenValue, size int) {
	evs.setValuesValues(values)
	evs.setSize(size)
}

func (evs *EigenValues) setValuesValues(values map[complex128]*EigenValue) {
	evs.values = values
}

func (evs *EigenValues) setSize(size int) {
	evs.size = size
}

func (evs *EigenValues) get(lambda complex128) *EigenValue {
	return evs.values[lambda]
}

func (evs *EigenValues) set(lambda complex128, ev *EigenValue) {
	evs.values[lambda] = ev
}

func (evs *EigenValues) setEigenValue(lambda complex128) {
	if ev := evs.get(lambda); ev != nil {
		ev.algebraicMultiplicity++
	} else {

	}
}

type EigenValue struct {
	lambda                complex128
	algebraicMultiplicity int
	geometricMultiplicity int
	vectors               *EigenVectors
}

func ConstructEigenValue(lambda complex128) *EigenValue {
	eigenValue := &EigenValue{lambda: lambda}
	return eigenValue
}

func (ev *EigenValue) Construct(lambda complex128) *EigenValue {
	return ev
}

func (ev *EigenValue) setSelf(evt *EigenValue) {

}

func (ev *EigenValue) setValues(lambda complex128, algebraicMultiplicity, geometricMultiplicity int, vectors *EigenVectors) {

}

func (ev *EigenValue) setLambda(lambda complex128) {
	ev.lambda = lambda
}

func (ev *EigenValue) addAlgebraicMultiplicity() {
	ev.algebraicMultiplicity++
}

type EigenVectors struct {
	VectorGroup
	geometricMultiplicity int
}
