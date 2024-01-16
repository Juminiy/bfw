package gmp

import "runtime"

func CallGCDirectly() {

	_slice := make([]byte, 0)

	for i := 0; i < 1<<32; i++ {
		_slice = append(_slice, 0)
	}

	_slice = nil

	runtime.GC()

}
