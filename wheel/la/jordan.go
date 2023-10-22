package la

type JordanMatrix struct {
	block     []*JordanBlock
	blockSize int
	size      int
}

func ConstructJordanMatrix() *JordanMatrix {
	jm := &JordanMatrix{}
	return jm.Construct()
}

func (jm *JordanMatrix) Construct() *JordanMatrix {
	return &JordanMatrix{}
}

// Sort
// order = true, lambda asc
// order = false, lambda desc
func (jm *JordanMatrix) Sort(order ...bool) {

}

func (jm *JordanMatrix) Display() *JordanMatrix {
	return jm
}

type JordanBlock struct {
	lambda       float64
	multiplicity int
}

func (jb *JordanBlock) Power() *Matrix {
	return &Matrix{}
}

func (jb *JordanBlock) Display() *JordanBlock {
	return jb
}
