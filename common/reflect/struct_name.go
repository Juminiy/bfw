package reflect

import "errors"

const (
	structNameNoSize int = 0
)

var (
	structNameIndexOutOfBound = errors.New("structName index is out of bound")
	structNameInvalidError    = errors.New("structName is invalid")
	NullStructName            = &StructName{}
)

type StructName struct {
	intField1         int
	intField2         int
	float64Field1     float64
	float64Field2     float64
	complex128Field1  complex128
	complex128Field2  complex128
	boolField1        bool
	boolField2        bool
	stringField1      string
	stringField2      string
	byteFiled1        byte
	byteFiled2        byte
	runeField1        rune
	runeField2        rune
	intSlice1d        []int
	intSlice2d        [][]int
	float64Slice1d    []float64
	float64Slice2d    [][]float64
	complex128Slice1d []complex128
	complex128Slice2d [][]complex128
	stringSlice1d     []string
	stringSlice2d     []string
}

func ConstructStructName() *StructName {
	structName := &StructName{}
	return structName.Construct()
}

func (sn *StructName) Construct() *StructName {
	return sn
}

func (sn *StructName) validate() bool {
	return true
}

func (sn *StructName) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !sn.validateOneIndex(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

func (sn *StructName) validateOneIndex(index int) bool {
	if index < 0 {
		return false
	}
	return true
}

func (sn *StructName) null() *StructName {
	return &StructName{}
}

func (sn *StructName) isNull() bool {
	return !sn.validate()
}

func (sn *StructName) setNull() {
	sn.setValues()
}

func (sn *StructName) makeCopy() *StructName {
	snCopy := &StructName{}
	snCopy.setValues()
	return snCopy
}

func (sn *StructName) assign() {
	sn.setValues()
}

func (sn *StructName) swap(snt *StructName) {
	snTemp := &StructName{}
	snTemp.setSelf(sn)
	sn.setSelf(snt)
	snt.setSelf(snTemp)
}

func (sn *StructName) getSelf() *StructName {
	return sn
}

func (sn *StructName) setSelf(snt *StructName) {
	sn.setValues()
}

func (sn *StructName) setValues() {

}

// ...
// reflect by go source code to generate source code
// do not consider deep copy
// func (sn *StructName) setFieldName(fieldName fieldType) {
//	sn.fieldName = fieldName
//}
// func (sn *StructName) GetFieldName() fieldType {
//	return sn.fieldName
//}

func (sn *StructName) Equal(snt *StructName) bool {
	return false
}
