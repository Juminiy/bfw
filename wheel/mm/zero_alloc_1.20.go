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

func GetAddressByBytes(a []byte) uintptr {
	return uintptr(unsafe.Pointer(unsafe.SliceData(a)))
}

func Int2Int64(a int) int64 {
	return *(*int64)(unsafe.Pointer(&a))
}

func Int642Int(a int64) int {
	return *(*int)(unsafe.Pointer(&a))
}

func GetAddressByInt64(a int64) uintptr {
	return uintptr(unsafe.Pointer(&a))
}

func GetAddressByInt(a int) uintptr {
	return uintptr(unsafe.Pointer(&a))
}
