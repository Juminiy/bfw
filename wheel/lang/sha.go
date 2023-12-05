package lang

import (
	"crypto/sha256"
	"math/bits"
)

func CountSHA256DiffBits(x, y []byte) int {
	xSha256 := sha256.Sum256(x)
	ySha256 := sha256.Sum256(y)
	diffCnt := 0

	// count 1 byte -> 4bit
	c0Byte := func(bt0 byte) uint8 {
		if bt0 >= '0' && bt0 <= '9' {
			bt0 -= '0'
		}
		if bt0 >= 'a' && bt0 <= 'f' {
			bt0 -= 'a'
		}
		if bt0 >= 'A' && bt0 <= 'F' {
			bt0 -= 'A'
		}
		return bt0
	}

	// count 2 byte -> 8bit
	x0Byte := func(xByte1, xByte2 byte) uint8 {
		hB := c0Byte(xByte1)
		lB := c0Byte(xByte2)
		return hB<<4 + lB
	}

	for i := 0; i < 32; i += 2 {
		x0B := x0Byte(xSha256[i], xSha256[i+1])
		y0B := x0Byte(ySha256[i], ySha256[i+1])
		diffCnt += bits.OnesCount8(x0B ^ y0B)
	}
	return diffCnt
}
