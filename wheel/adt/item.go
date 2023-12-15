package adt

import "errors"

var (
	infinityCanNotZeroError = errors.New("infinity cannot be zero")
	nInf                    = NInf{}
	pInf                    = PInf{}
)

type Item interface {
	Less(Item) bool
}

type NInf struct{}

func (NInf) Less(Item) bool {
	return true
}

type PInf struct{}

func (PInf) Less(Item) bool {
	return false
}

func GetInf(sign int) Item {
	if sign == 0 {
		panic(infinityCanNotZeroError)
	} else if sign < 0 {
		return nInf
	} else {
		return pInf
	}
}

func less(a, b Item) bool {
	if a == nInf ||
		b == pInf {
		return true
	} else if a == pInf ||
		b == nInf {
		return false
	}
	return a.Less(b)
}

func ItemLess(a, b Item) bool {
	return less(a, b)
}

type Int int

func (int Int) Less(than Item) bool {
	if than == nInf {
		return false
	} else if than == pInf {
		return true
	} else {
		thanInt := than.(Int)
		return int < thanInt
	}
}
