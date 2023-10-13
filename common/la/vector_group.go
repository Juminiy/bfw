package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
)

const (
	vectorGroupNoSize          int = 0
	vectorGroupIndexOutOfBound int = -1
)

var (
	vectorGroupIndexOutOfBoundError = errors.New("vector group index is out of bound")
	vectorGroupCanNotSchmidtError   = errors.New("vector group cannot schmidt")
	vectorGroupInvalidError         = errors.New("vector group is invalid")
	NullVectorGroup                 = &VectorGroup{}
)

// VectorGroup
// shape = true -> line vector group
// shape = false -> row vector group
type VectorGroup struct {
	group []*Vector
	size  int
	shape bool
}

func (vg *VectorGroup) validate() (bool, int) {
	if vg.size == vectorGroupNoSize ||
		vg.group == nil ||
		len(vg.group) != vg.size {
		return false, vectorGroupNoSize
	} else {
		vSize := vg.group[0].size
		for idx := 0; idx < vg.size; idx++ {
			if !vg.group[idx].validate() ||
				vg.group[idx].size != vSize {
				return false, vectorGroupNoSize
			}
		}
		return true, vSize
	}
}

func (vg *VectorGroup) validateOneIndex(index, vSize int) bool {
	if index < 0 ||
		index >= vg.size ||
		index >= vSize {
		return false
	}
	return true
}

func (vg *VectorGroup) validateIndex(index ...int) bool {
	ok, vSize := vg.validate()
	if !ok {
		panic(vectorGroupInvalidError)
	}
	if indexLen := len(index); indexLen > 0 {
		if indexLen >= 1 {
			if index[0] < 0 ||
				index[0] >= vg.size {
				return false
			}
		} else if indexLen >= 2 {
			if index[1] < 0 ||
				index[0] >= vSize {
				return false
			}
		} else {
			for indexIdx := 2; indexIdx < indexLen; indexIdx++ {
				if !vg.validateOneIndex(index[indexIdx], vSize) {
					return false
				}
			}
		}
	}
	return true
}

func (vg *VectorGroup) assign(size, vSize int) {
	vg.setValues(make([]*Vector, size), size, true)
	for idx := 0; idx < size; idx++ {
		vg.group[idx] = &Vector{}
		vg.group[idx].setValues(make([]float64, vSize), vSize, true)
	}
}

func (vg *VectorGroup) set(index int, vector *Vector) {
	if index < 0 ||
		index >= vg.size {
		panic(vectorGroupIndexOutOfBoundError)
	}
	vg.group[index] = vector
}

func (vg *VectorGroup) get(index int) *Vector {
	if index < 0 ||
		index >= vg.size {
		panic(vectorGroupIndexOutOfBoundError)
	}
	return vg.group[index]
}

func (vg *VectorGroup) setElem(index, vIndex int, value float64) {
	if !vg.validateIndex(index, vIndex) {
		return
	}
	vg.group[index].set(vIndex, vIndex, value)
}

func (vg *VectorGroup) getValue(index, vIndex int) float64 {
	return vg.group[index].get(vIndex, vIndex)
}
func (vg *VectorGroup) setGroup(group []*Vector) {
	vg.group = group
}

func (vg *VectorGroup) setSize(size int) {
	vg.size = size
}

func (vg *VectorGroup) setShape(shape bool) {
	vg.shape = shape
}

func (vg *VectorGroup) setValues(group []*Vector, size int, shape bool) {
	vg.setGroup(group)
	vg.setSize(size)
	vg.setShape(shape)
}

func (vg *VectorGroup) setSlice(slice [][]float64) {
	for rowIdx, row := range slice {
		for lineIdx, ele := range row {
			vg.setElem(rowIdx, lineIdx, ele)
		}
	}
}

func (vg *VectorGroup) setSelf(vgt *VectorGroup) {
	vg.setValues(vgt.group, vgt.size, vgt.shape)
}

func (vg *VectorGroup) getVSize() int {
	ok, vSize := vg.validate()
	if !ok {
		panic(vectorGroupInvalidError)
	}
	return vSize
}

func (vg *VectorGroup) makeCopy() *VectorGroup {
	vgCopy := &VectorGroup{}
	vgCopy.setValues(make([]*Vector, vg.size), vg.size, vg.shape)
	for idx := 0; idx < vgCopy.size; idx++ {
		vgCopy.set(idx, vg.get(idx).makeCopy())
	}
	return vgCopy
}

func (vg *VectorGroup) convertToMatrix() *Matrix {
	ok, vSize := vg.validate()
	if !ok {
		panic(vectorGroupInvalidError)
	}
	matrix := &Matrix{}
	if vg.shape {
		matrix.assign(vSize, vg.size)
		for lineIdx := 0; lineIdx < vSize; lineIdx++ {
			matrix.setLine(lineIdx, vg.get(lineIdx).slice)
		}
	} else {
		matrix.setValues(make([][]float64, vg.size), vg.size, vSize)
		for idx := 0; idx < vg.size; idx++ {
			matrix.setRow(idx, make([]float64, vSize))
			copy(matrix.slice[idx], vg.get(idx).slice)
		}
	}
	return matrix
}

func (vg *VectorGroup) Construct(real2DArray [][]float64, shape ...bool) *VectorGroup {
	if real2DArray == nil ||
		len(real2DArray) == vectorGroupNoSize {
		return vg
	}
	var (
		size   int  = len(real2DArray)
		vSize  int  = len(real2DArray[0])
		vShape bool = true
	)
	if len(shape) > 0 && !shape[0] {
		vShape = shape[0]
	}
	for idx := 0; idx < size; idx++ {
		vSize = lang.MaxInt(vSize, len(real2DArray[idx]))
	}
	vg.setSize(size)
	vg.setShape(vShape)
	// without padding

	// 1. set each element
	//vg.assign(size, vSize)
	//vg.setSlice(real2DArray)

	// 2. set each vector
	vg.setGroup(make([]*Vector, size))
	for idx := 0; idx < size; idx++ {
		vg.set(idx, ConstructVector(lang.GetReal2DArrayLine(real2DArray, idx), vShape))
	}
	return vg
}

func ConstructVectorGroup(real2DArray [][]float64, shape ...bool) *VectorGroup {
	vg := &VectorGroup{}
	return vg.Construct(real2DArray, shape...)
}

func (vg *VectorGroup) isNull() bool {
	ok, _ := vg.validate()
	return !ok
}

func (vg *VectorGroup) null() *VectorGroup {
	return &VectorGroup{}
}

func (vg *VectorGroup) setNull() {
	vg.setValues(nil, vectorGroupNoSize, false)
}

func (vg *VectorGroup) Matrix() *Matrix {
	return vg.convertToMatrix()
}

// schmidt
// consider whether to expand the vector space
// current version not
func (vg *VectorGroup) schmidt() *VectorGroup {
	if vg.shape &&
		vg.size > 0 {
		vgBeta := vg.makeCopy()
		for groupIdx := 1; groupIdx < vgBeta.size; groupIdx++ {
			for projectionIdx := 0; projectionIdx < groupIdx; projectionIdx++ {
				vgBeta.group[groupIdx].Sub(
					vgBeta.get(projectionIdx).
						MulLambda(vgBeta.get(projectionIdx).DotMul(vg.get(groupIdx)) /
							vgBeta.get(projectionIdx).InnerMul()))
			}
		}
		return vgBeta
	}
	panic(vectorGroupCanNotSchmidtError)
}

func (vg *VectorGroup) GetSchmidt() *VectorGroup {
	vgCopy := vg.makeCopy()
	return vgCopy.schmidt()
}

func (vg *VectorGroup) unit() *VectorGroup {
	for idx := 0; idx < vg.size; idx++ {
		vg.set(idx, vg.get(idx).GetUnit())
	}
	return vg
}

func (vg *VectorGroup) GetUnit() *VectorGroup {
	vgCopy := vg.makeCopy()
	return vgCopy.unit()
}

func (vg *VectorGroup) validateUnit() bool {
	for idx := 0; idx < vg.size; idx++ {
		if !vg.get(idx).validateUnit() {
			return false
		}
	}
	return true
}

func (vg *VectorGroup) validateOrthonormal() bool {
	for idxI := 0; idxI < vg.size; idxI++ {
		for idxJ := idxI + 1; idxJ < vg.size; idxJ++ {
			if !vg.get(idxI).validateOrthogonal(vg.get(idxJ)) {
				return false
			}
		}
	}
	return true
}

func (vg *VectorGroup) ValidateOrthonormalBasis() bool {
	return vg.validateUnit() &&
		vg.validateOrthonormal()
}

func (vg *VectorGroup) Display() *VectorGroup {
	if vg.isNull() {
		fmt.Println("[null]")
		return vg
	}
	vg.Matrix().Display()
	return vg
}

func (vg *VectorGroup) IsBasis() bool {
	return false
}

func (vg *VectorGroup) IsLinearCorrelation() bool {
	return false
}
