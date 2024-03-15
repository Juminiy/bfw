package fp

func curryingAddInt(e ...int) func(...int) int {
	return func(i ...int) int {
		return i[0]
	}
}

type func1Ptr func(...int) int

func curryingAdd3Int(a, b, c int) func1Ptr {
	var f1 func1Ptr
	return f1
}
