package generic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneric(t *testing.T) {
	t.Run("type ref", func(t *testing.T) {
		ref0Type := reflect.TypeOf(Ref0{})
		ref1Type := reflect.TypeOf(Ref1{})
		ref2Type := reflect.TypeOf(Ref2{})
		ref3Type := reflect.TypeOf(Ref3{})
		fmt.Println(ref0Type, "\n", ref1Type, "\n", ref2Type)
		fmt.Println(ref0Type == ref1Type)
		fmt.Println(ref1Type == ref2Type)
		fmt.Println(ref3Type == ref0Type)
	})
}
