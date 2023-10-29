package la

type ComplexMatrix struct {
	slice [][]complex128
}

func (cm *ComplexMatrix) makeCopy() *ComplexMatrix {
	return cm
}

// Aᵀ
func (cm *ComplexMatrix) transpose() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) conjugate() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) CTranspose() *ComplexMatrix {
	return cm
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

// UnitarySimilar
// return UᴴAU = U⁻¹AU (UᴴU = UUᴴ = I)
func (cm *ComplexMatrix) UnitarySimilar(U *UnitaryMatrix) *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) IsUnitarySimilar(B *ComplexMatrix) bool {
	return false
}

func (cm *ComplexMatrix) Matrix() *Matrix {
	return cm.convertToMatrix()
}
