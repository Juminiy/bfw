package la

const (
	identityMatrixValue float64 = 1.0
)

var (
	NullIdentity = &Identity{}
)

type Identity struct {
	size int
}

func (identity *Identity) null() *Identity {
	return &Identity{}
}

func (identity *Identity) Matrix() *Matrix {
	matrix := &Matrix{}
	matrix.assign(identity.size, identity.size)
	for rowIdx := 0; rowIdx < matrix.GetRowSize(); rowIdx++ {
		for lineIdx := 0; lineIdx < matrix.GetLineSize(); lineIdx++ {
			if rowIdx == lineIdx {
				matrix.set(rowIdx, lineIdx, 1)
			} else {
				matrix.set(rowIdx, lineIdx, 0)
			}
		}
	}
	return matrix
}

func (identity *Identity) Mul(id *Identity) *Identity {
	if id != nil &&
		identity.size == id.size {
		return identity
	}
	panic(matrixCanNotMultiplyError)
}

func (identity *Identity) MulLambda(lambda float64) *Matrix {
	return identity.Matrix().MulLambda(lambda)
}

func (identity *Identity) Display() *Identity {
	return identity.Matrix().Display().Identity()
}
