package la

type EigenValues struct {
	size   int
	values []*EigenValue
}

type EigenValue struct {
	lambda                float64
	algebraicMultiplicity int
	geometricMultiplicity int
	vectors               *EigenVector
}

type EigenVector struct {
	VectorGroup
	geometricMultiplicity int
}
