package la

var (
	NullJordanMatrix = &JordanMatrix{}
)

type JordanMatrix struct {
	block     []*JordanBlock
	blockSize int
	size      int
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
	return NullMatrix
}

func (jb *JordanBlock) Display() *JordanBlock {
	return jb
}
