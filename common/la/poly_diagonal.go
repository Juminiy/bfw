package la

var (
	NullPolyDiagonal = &PolyDiagonal{}
)

type PolyDiagonal struct {
	slice []*Poly
	size  int
}
