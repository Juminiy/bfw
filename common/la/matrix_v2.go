package la

var (
	NullGenericMatrix = &GenericMatrix{}
)

type GenericMatrix struct {
	slice    [][]complex128
	rowSize  int64
	lineSize int64
}

func (gm *GenericMatrix) makeCopy() *GenericMatrix {
	return gm
}

func (gm *GenericMatrix) transpose() *GenericMatrix {
	return gm
}

func (gm *GenericMatrix) GetTranspose() *GenericMatrix {
	gmCopy := gm.makeCopy()
	return gmCopy.transpose()
}

func (gm *GenericMatrix) conjugate() *GenericMatrix {
	return gm
}

func (gm *GenericMatrix) GetConjugate() *GenericMatrix {
	gmCopy := gm.makeCopy()
	return gmCopy.conjugate()
}

func (gm *GenericMatrix) GetConjugateTranspose() *GenericMatrix {
	gmCopy := gm.makeCopy()
	return gmCopy.conjugate().transpose()
}

func (gm *GenericMatrix) null() *GenericMatrix {
	return &GenericMatrix{}
}

func (gm *GenericMatrix) convertToMatrix() *Matrix {
	return &Matrix{}
}

func (gm *GenericMatrix) Matrix() *Matrix {
	return gm.convertToMatrix()
}
