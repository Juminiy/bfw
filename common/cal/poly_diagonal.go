package cal

var (
	NullPolyDiagonal = &PolyDiagonal{}
)

type PolyDiagonal struct {
	slice []*Poly
	size  int
}
