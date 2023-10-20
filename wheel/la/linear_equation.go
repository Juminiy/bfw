package la

// MultipleVariablesLinearEquation
// to solve X
// loc = true,  A*X = B
// loc = false, X*A = B
type MultipleVariablesLinearEquation struct {
	A   *Matrix
	B   *VectorGroup
	loc bool
}

type MVLE MultipleVariablesLinearEquation

func (mvle *MVLE) Solve() *VectorGroup {
	if mvle.loc {
		return mvle.solveL()
	} else {
		return mvle.solveR()
	}
}

func (mvle *MVLE) solveL() *VectorGroup {
	return &VectorGroup{}
}

func (mvle *MVLE) solveR() *VectorGroup {
	return &VectorGroup{}
}
