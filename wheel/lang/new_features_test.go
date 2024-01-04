package lang

import (
	"testing"
)

func TestNewFeaturesClear(t *testing.T) {
	NewFeaturesClear()
}

func TestPointerNew(t *testing.T) {
	PointerNew()
}

func TestGetGenericsZeroValue(t *testing.T) {

}

func TestAReceiver_Inc(t *testing.T) {
	a := AReceiver{}
	a.Inc()
	a.Print()
	a.Inc()
	a.Print()
	a.Plus(1)
	a.Print()
	a.Plus(1)
	a.Print()
	//gs := GenericStruct[int]{}
	//fmt.Println(gs)
}
