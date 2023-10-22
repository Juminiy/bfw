package la

type ComplexMatrix struct {
	slice      [][]complex128
	rowSize    int64
	columnSize int64
}

func (cm *ComplexMatrix) makeCopy() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) transpose() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) GetTranspose() *ComplexMatrix {
	gmCopy := cm.makeCopy()
	return gmCopy.transpose()
}

func (cm *ComplexMatrix) conjugate() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) GetConjugate() *ComplexMatrix {
	gmCopy := cm.makeCopy()
	return gmCopy.conjugate()
}

func (cm *ComplexMatrix) CTranspose() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) GetConjugateTranspose() *ComplexMatrix {
	gmCopy := cm.makeCopy()
	return gmCopy.conjugate().transpose()
}

func (cm *ComplexMatrix) GetPhaseAngle() *Matrix {
	return cm.phaseAngle()
}

func (cm *ComplexMatrix) phaseAngle() *Matrix {
	return &Matrix{}
}

func (cm *ComplexMatrix) null() *ComplexMatrix {
	return &ComplexMatrix{}
}

func (cm *ComplexMatrix) convertToMatrix() *Matrix {
	return &Matrix{}
}

func (cm *ComplexMatrix) Matrix() *Matrix {
	return cm.convertToMatrix()
}
