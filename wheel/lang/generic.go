package lang

import "cmp"

type UndefinedStruct struct {
	FieldIntA     int
	FiledFloat64B float64
	FieldStringC  string
}

// Compiler Auto Infer

func OrderedAdd[T cmp.Ordered](a, b T) T {
	return a + b
}
