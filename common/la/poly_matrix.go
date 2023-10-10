package la

var (
	NullPolyMatrix = &PolyMatrix{}
)

type PolyMatrix struct {
	slice       [][]*Poly
	rowSize     int
	lineSize    int
	coefficient float64
}

func (pm *PolyMatrix) det() *Poly {
	return &Poly{}
}

func (pm *PolyMatrix) Det() *Poly {
	return pm.det()
}

func (pm *PolyMatrix) Equal() bool {
	return false
}

func (pm *PolyMatrix) eigenMatrixRowExchangeET() {

}

func (pm *PolyMatrix) eigenMatrixLineExchangeET() {

}

func (pm *PolyMatrix) eigenMatrixRowMulLambdaET() {

}

func (pm *PolyMatrix) eigenMatrixLineMulLambdaET() {

}

func (pm *PolyMatrix) eigenMatrixRow1MulPolyAddRow2ET() {

}

func (pm *PolyMatrix) eigenMatrixLine1MulPolyAddLine2ET() {

}

func (pm *PolyMatrix) eigenMatrixElementaryTransformation() {

}

func (pm *PolyMatrix) Smith() {

}

func (pm *PolyMatrix) greatestCommonFactor() {

}

func (pm *PolyMatrix) D(k int) {

}

func (pm *PolyMatrix) d(k int) {

}

func (pm *PolyMatrix) Display() {

}
