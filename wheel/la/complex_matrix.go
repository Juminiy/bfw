package la

type ComplexMatrix struct {
	slice      [][]complex128
	rowSize    int64
	columnSize int64
}

func (gm *ComplexMatrix) makeCopy() *ComplexMatrix {
	return gm
}

func (gm *ComplexMatrix) transpose() *ComplexMatrix {
	return gm
}

func (gm *ComplexMatrix) GetTranspose() *ComplexMatrix {
	gmCopy := gm.makeCopy()
	return gmCopy.transpose()
}

func (gm *ComplexMatrix) conjugate() *ComplexMatrix {
	return gm
}

func (gm *ComplexMatrix) GetConjugate() *ComplexMatrix {
	gmCopy := gm.makeCopy()
	return gmCopy.conjugate()
}

func (gm *ComplexMatrix) GetConjugateTranspose() *ComplexMatrix {
	gmCopy := gm.makeCopy()
	return gmCopy.conjugate().transpose()
}

func (gm *ComplexMatrix) GetPhaseAngle() *Matrix {
	return gm.phaseAngle()
}

func (gm *ComplexMatrix) phaseAngle() *Matrix {
	return &Matrix{}
}

func (gm *ComplexMatrix) null() *ComplexMatrix {
	return &ComplexMatrix{}
}

func (gm *ComplexMatrix) convertToMatrix() *Matrix {
	return &Matrix{}
}

func (gm *ComplexMatrix) Matrix() *Matrix {
	return gm.convertToMatrix()
}
