package learn

import (
	"errors"
	"reflect"
)

type Neko struct{}

func ReadAny2NekoPtr(a any) (*Neko, error) {
StructOrPointer:
	aType := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)
	switch aType.Kind() {
	case reflect.Struct:
		if aType.String() == "learn.Neko" {
			aNeko := a.(Neko)
			return &aNeko, nil
		} else {
			return nil, errors.New("not a Neko Type")
		}
	case reflect.Pointer:
		//a = aValue.Elem()
		//goto StructOrPointer
		if aTypeString := aType.String(); aTypeString == "**learn.Neko" {
			a = *(a.(**Neko))
			return a.(*Neko), nil
		} else if aTypeString == "*learn.Neko" {
			return a.(*Neko), nil
		} else {
			return nil, errors.New("not a Neko Pointer")
		}
		// if has more, need to iterator dereference
	case reflect.Array, reflect.Slice:
		aType = aType.Elem()
		if aValue.Len() == 0 {
			a = reflect.New(aType).Interface()
		} else {
			a = aValue.Index(0).Interface()
		}

		goto StructOrPointer
	default:
		return nil, errors.New("not a Neko struct")
	}
}

func GetNone() any {
	type None struct {
		A int
		B string
		C complex128
	}
	return new(None)
}
