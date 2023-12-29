package lang

import (
	"fmt"
	"testing"
)

func TestNewFeaturesClear(t *testing.T) {
	NewFeaturesClear()
}

func TestPointerNew(t *testing.T) {
	PointerNew()
}

func TestGetGenericsZeroValue(t *testing.T) {
	gs := GenericStruct{}
	fmt.Println(gs)
}
