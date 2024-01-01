//go:build !go1.20

package mm

import (
	"unsafe"
)

// zero memory allocation for String2Bytes or Bytes2String type convert

func String2Bytes(a string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Size int
	}{a, len(a)}))
}

func Bytes2String(a []byte) string {
	return *(*string)(unsafe.Pointer(&a))
}

func GetAddressByString(a string) uintptr {
	return uintptr(unsafe.Pointer(&struct {
		string
	}{a}))
}
