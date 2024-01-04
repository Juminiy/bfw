package mm

import (
	"fmt"
	"testing"
)

func TestBytes2String(t *testing.T) {
	var ByteSlice []byte
	for i := 0; i < (1 << 8); i++ {
		ByteSlice = append(ByteSlice, byte(i))
	}
	fmt.Printf("%p\n", ByteSlice)
	fmt.Printf("0x%x\n", GetAddressByString(string(ByteSlice)))
	fmt.Printf("0x%x\n", GetAddressByString(Bytes2String(ByteSlice)))
}

func TestInt642Int(t *testing.T) {
	var (
		i64 int64 = 0
		//i   int
	)
	fmt.Printf("%p\n", &i64)
	fmt.Printf("0x%x\n", GetAddressByInt64(i64))
	fmt.Printf("0x%x\n", GetAddressByInt(int(i64)))
}
