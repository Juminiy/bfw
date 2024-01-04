package learn

import "fmt"

type UndefinedInterface interface{}

type UndefinedStruct struct{}

type DefinedInterface interface {
	Def(int) int
}

type DefinedStruct struct {
	FieldIntOuterA int
	fieldIntInnerA int
}

func (ds *DefinedStruct) Inc() *DefinedStruct {
	ds.fieldIntInnerA++
	return ds
}

func (ds *DefinedStruct) Dec() *DefinedStruct {
	ds.fieldIntInnerA--
	return ds
}

func (ds *DefinedStruct) Print() *DefinedStruct {
	fmt.Println(ds.fieldIntInnerA)
	return ds
}

func (ds *DefinedStruct) Def(a int) int {
	fmt.Println(ds.fieldIntInnerA)
	return ds.fieldIntInnerA
}
