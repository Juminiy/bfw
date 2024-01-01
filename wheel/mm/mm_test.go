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
