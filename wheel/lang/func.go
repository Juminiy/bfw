package lang

import "fmt"

func dup(a any) {
	fmt.Printf("args:%p\nassertfunc:%p\n", a, a.(func()))
}
