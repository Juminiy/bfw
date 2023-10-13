package la

type MatrixPoly struct {
	matrix *Matrix
	poly   *Poly
}

func (mp *MatrixPoly) Poly() *Poly {
	return &Poly{}
}
