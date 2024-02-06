package runtime

//go:generate stringer -type Pullgy -output pulggy.go
type Pulggy struct {
	Arc  int
	Bnmi string
	Cone *Pulggy
}

type Rem struct {
	Pillo Pulggy
}

func (r *Rem) DoRem() {
	r.Pillo.Cone = &r.Pillo
	x := Pulggy{}
	r.Pillo = x
}

var (
	GlobalPulggy Pulggy
)

const (
// EP Pulggy = Pulggy{}
)
