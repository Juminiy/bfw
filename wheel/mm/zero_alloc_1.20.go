//go:build go1.20

package mm

import "unsafe"

// zero memory allocation for String2Bytes or Bytes2String type convert

func String2Bytes(a string) []byte {
	return unsafe.Slice(unsafe.StringData(a), len(a))
}

func Bytes2String(a []byte) string {
	return unsafe.String(unsafe.SliceData(a), len(a))
}

func GetAddressByString(a string) uintptr {
	return uintptr(unsafe.Pointer(unsafe.StringData(a)))
}
