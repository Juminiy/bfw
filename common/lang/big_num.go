package lang

import (
	"errors"
	"fmt"
)

const (
	spaceByte        byte = ' '
	baseByteBin      byte = 'B'
	baseByteOct      byte = 'O'
	baseByteDec      byte = 'D'
	baseByteHex      byte = 'H'
	signBytePositive byte = '+'
	signByteNegative byte = '-'
	numByteZero      byte = '0'
	numByteOne       byte = '1'
	numByteTwo       byte = '2'
	numByteThree     byte = '3'
	numByteFour      byte = '4'
	numByteFive      byte = '5'
	numByteSix       byte = '6'
	numByteSeven     byte = '7'
	numByteEight     byte = '8'
	numByteNine      byte = '9'
	numByteTen       byte = 'a'
	numByteEleven    byte = 'b'
	numByteTwelve    byte = 'c'
	numByteThirteen  byte = 'd'
	numByteFourteen  byte = 'e'
	numByteFifteen   byte = 'f'
	eqByte           byte = '='
	gtByte           byte = '>'
	lsByte           byte = '<'
	bigNumNoSize     int  = 0
)

var (
	bigNumInvalidError         = errors.New("big num is invalid")
	bigNumIndexOutOfBoundError = errors.New("big num index is out of bound")
	inputNumStringInvalidError = errors.New("big num input num string is invalid")
	notSupporttedBaseError     = errors.New("has not support the base")
)

type BigNum struct {
	slice []byte
	size  int
	sign  byte
	base  byte
}

func ConstructBigNum(num string, base ...int) *BigNum {
	bn := &BigNum{}
	return bn.Construct(num, base...)
}

// Construct
// "123456789"
// "-123456789"
// " 123456789 "
// " -123456789 "
func (bn *BigNum) Construct(num string, base ...int) *BigNum {
	num = TruncateStringPrefixSuffixSpace(num)
	baseByte := baseByteDec
	signByte := signBytePositive
	if len(base) > 0 {
		baseByte = getBase(base[0])
	}
	if num[0] == signByteNegative {
		num = num[1:]
		signByte = signByteNegative
	}
	bn.setValues(make([]byte, len(num)), len(num), signByte, baseByte)
	for idx := 0; idx < bn.size; idx++ {
		char := num[idx]
		if !validateNumByte(char) {
			panic(inputNumStringInvalidError)
		}
		bn.setElem(idx, char)
	}
	return bn
}

func getBase(base int) byte {
	switch base {
	case 2:
		{
			return baseByteBin
		}
	case 8:
		{
			return baseByteOct
		}
	case 10:
		{
			return baseByteDec
		}
	case 16:
		{
			return baseByteHex
		}
	default:
		{
			panic(notSupporttedBaseError)
		}
	}
	return 0
}

func (bn *BigNum) validateOneIndex(index int) bool {
	return index >= 0 && index < bn.size
}

func (bn *BigNum) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !bn.validateOneIndex(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

func validateNumByte(char byte) bool {
	return char >= numByteZero &&
		char <= numByteFifteen
}

func (bn *BigNum) validateElem(index int) bool {
	return validateNumByte(bn.getElem(index))
}

func (bn *BigNum) validateSlice() bool {
	for idx := 0; idx < bn.size; idx++ {
		if !bn.validateElem(idx) {
			return false
		}
	}
	return true
}

func (bn *BigNum) validateSize() bool {
	return bn.size != bigNumNoSize &&
		len(bn.slice) == bn.size
}

func (bn *BigNum) validateSign() bool {
	return bn.sign == signByteNegative ||
		bn.sign == signBytePositive ||
		bn.sign == spaceByte
}

func (bn *BigNum) validateBase() bool {
	return bn.base == baseByteBin ||
		bn.base == baseByteOct ||
		bn.base == baseByteDec ||
		bn.base == baseByteHex
}

func (bn *BigNum) validate() bool {
	return bn.validateSlice() &&
		bn.validateSize() &&
		bn.validateSign() &&
		bn.validateBase()
}

func (bn *BigNum) makeCopy() *BigNum {
	bnCopy := &BigNum{}
	bnCopy.setValues(make([]byte, bn.size), bn.size, bn.sign, bn.base)
	copy(bnCopy.slice, bn.slice)
	return bnCopy
}

func (bn *BigNum) assign(size int) {
	bn.setValues(make([]byte, size), size, signBytePositive, baseByteDec)
}

func (bn *BigNum) setSelf(bnt *BigNum) {
	bn.setValues(bnt.slice, bnt.size, bnt.sign, bnt.base)
}

func (bn *BigNum) getSelf() *BigNum {
	return bn
}

func (bn *BigNum) setValues(slice []byte, size int, sign, base byte) {
	bn.setSlice(slice)
	bn.setSize(size)
	bn.setSign(sign)
	bn.setBase(base)
}

func (bn *BigNum) setSlice(slice []byte) {
	bn.slice = slice
}

func (bn *BigNum) setSize(size int) {
	bn.size = size
}

func (bn *BigNum) setSign(sign byte) {
	bn.sign = sign
}

func (bn *BigNum) setBase(base byte) {
	bn.base = base
}

func (bn *BigNum) getElem(index int) byte {
	if !bn.validateIndex(index) {
		panic(bigNumIndexOutOfBoundError)
	}
	return bn.slice[index]
}

func (bn *BigNum) getSetElem(index int) byte {
	if !bn.validateIndex(index) {
		panic(bigNumIndexOutOfBoundError)
	}
	if !validateNumByte(bn.slice[index]) {
		bn.slice[index] = numByteZero
	}
	return bn.slice[index]
}

func (bn *BigNum) setElem(index int, char byte) {
	if !bn.validateIndex(index) {
		panic(bigNumIndexOutOfBoundError)
	}
	bn.slice[index] = char
}

func (bn *BigNum) appendSlice(slice ...byte) {
	bn.slice = append(bn.slice, slice...)
	bn.size += len(slice)
}

func (bn *BigNum) resetSize(size int) {
	if bn.size > size {
		bn.setSlice(bn.slice[:size])
	} else if bn.size < size {
		appendSlice := make([]byte, size-bn.size)
		bn.appendSlice(appendSlice...)
	} else {

	}
	bn.setSize(size)
}

func (bn *BigNum) resetSign() {
	if bn.sign == signBytePositive {
		bn.setSign(signByteNegative)
	} else {
		bn.setSign(signBytePositive)
	}
}

// resetBase
// a bit of complex
func (bn *BigNum) resetBase() {

}

func (bn *BigNum) isZero() bool {
	return bn.validate() &&
		bn.size == 1 &&
		bn.getElem(0) == numByteZero
}

func (bn *BigNum) setZero() {
	bn.setSlice([]byte{'0'})
	bn.setSize(1)
}

func (bn *BigNum) Abs() *BigNum {
	bn.setSign(signBytePositive)
	return bn
}

func (bn *BigNum) Opp() *BigNum {
	bn.resetSign()
	return bn
}

func (bn *BigNum) Equal(bnt *BigNum) bool {
	if bnt == nil ||
		!bn.validate() ||
		!bnt.validate() {
		panic(bigNumInvalidError)
	}
	return bn.size == bnt.size &&
		string(bn.slice) == string(bnt.slice)
}

// compareTo
// 2 > 1
// 2 > -1
// -2 < -1
// -2 < 1

// 2 > -2
// 2 == 2
// 2 < -2

// 1 < 2
// 1 > -2
// -1 < 2
// -1 > -2
func (bn *BigNum) compareTo(bnt *BigNum) byte {
	if bnt == nil ||
		!bn.validate() ||
		!bnt.validate() {
		panic(bigNumInvalidError)
	}
	numRes := bn.compareNum(bnt)
	signRes := bn.compareSign(bnt)
	if numRes == gtByte {
		if signRes == gtByte ||
			signRes == eqByte {
			return gtByte
		} else {
			return lsByte
		}
	} else if numRes == eqByte {
		return signRes
	} else {
		if signRes == lsByte ||
			signRes == eqByte {
			return lsByte
		} else {
			return gtByte
		}
	}
}

// + > -
// + = +
// - < +
// - = -
func (bn *BigNum) compareSign(bnt *BigNum) byte {
	if bn.sign == bnt.sign {
		return eqByte
	}
	if bn.sign == signBytePositive {
		return gtByte
	} else {
		return lsByte
	}
}

func (bn *BigNum) compareNum(bnt *BigNum) byte {
	bnStr, bntStr := string(bn.slice), string(bnt.slice)
	if bnStr > bntStr {
		return gtByte
	} else if bnStr == bntStr {
		return eqByte
	} else {
		return lsByte
	}
}

func (bn *BigNum) prefixPadding(destSize int, char byte) {
	if destSize <= bn.size {
		return
	}
	slice := make([]byte, destSize-bn.size)
	for idx := 0; idx < len(slice); idx++ {
		slice[idx] = char
	}
	bn.setSlice(append(slice, bn.slice...))
	bn.setSize(destSize)
}

func (bn *BigNum) prefixZeroPadding(destSize int) {
	bn.prefixPadding(destSize, numByteZero)
}

func (bn *BigNum) setElemOpt(index int, opt rune, char ...byte) {
	switch opt {
	case '+':
		{
			sumNum := bn.getElem(index) - numByteZero + char[0] - numByteZero
			if sumNum >= 10 {
				bn.setElem(index-1, bn.getElem(index-1)+1)
			}
			sumNum %= 10
			bn.setElem(index, sumNum+numByteZero)
		}
	case '-':
		{
			num1, num2 := bn.getElem(index), char[0]
			if num1 >= num2 {
				num1 = num1 - num2 + numByteZero
			} else {
				num1 = num1 + 10 - num2 + numByteZero
				bn.setElem(index-1, bn.getElem(index-1)-1)
			}
			bn.setElem(index, num1)
		}
	case '*':
		{
			//c[i+j] += a[i]*b[j]
			//c[i+j+1] += c[i+j]/10
			//c[i+j] %=10;
			byteNum1, byteNum2, byteNum3 := bn.getSetElem(index), char[0], char[1]
			num1, num2, num3 := byteNum1-numByteZero, byteNum2-numByteZero, byteNum3-numByteZero
			num1 += num2 * num3
			bn.setElem(index-1, bn.getSetElem(index-1)+num1/10)
			bn.setElem(index, num1%10+numByteZero)
		}
	case '/':
		{

		}
	default:
		{

		}
	}
}

// one2OneOpt
// make sure the bn.compareNum(bnt) = gtByte
//
// 123456789
// 123456789
// 246913578

// 111111000
// 000000999
// 111110001

// 0999000
// 0000999
func (bn *BigNum) one2OneOpt(opt rune, bnt *BigNum) *BigNum {
	if bnt == nil ||
		!bn.validate() ||
		!bnt.validate() {
		return bn
	}
	for idx := bn.size - 1; idx > 0; idx-- {
		bn.setElemOpt(idx, opt, bnt.getElem(idx))
	}
	return bn
}

// Add
// +a add +b -> +(a+b)
// +a add -b -> if a=b, res=0; if a>b, res=+(a-b); if a<b, res=-(b-a)
// -a add +b -> if a=b, res=0; if a>b, res=-(a-b); if a<b, res=+(b-a)
// -a add -b -> -(a+b)
func (bn *BigNum) Add(bnt *BigNum) *BigNum {
	bnDestSize := MaxInt(bn.size, bnt.size) + 1
	bn.prefixZeroPadding(bnDestSize)
	bnt.prefixZeroPadding(bnDestSize)
	if signRes := bn.compareSign(bnt); signRes == eqByte {
		return bn.one2OneOpt('+', bnt)
	} else {
		numRes := bn.compareNum(bnt)
		if numRes == eqByte {
			bn.setZero()
			return bn
		} else if numRes == gtByte {
			return bn.one2OneOpt('-', bnt)
		} else {
			// whether to change bnt
			//bntCopy := bnt.makeCopy()
			bnt.one2OneOpt('-', bn)
			bn.setSelf(bnt)
			//bnt.setSelf(bntCopy)
			return bn
		}
	}
}

// Sub
// +a sub +b -> if a=b, res=0; if a>b, res=+(a-b); if a<b, res=-(b-a)
// +a sub -b -> +(a+b)
// -a sub +b -> -(a+b)
// -a sub -b -> if a=b, res=0; if a>b, res=-(a-b); if a<b, res=+(b-a)
func (bn *BigNum) Sub(bnt *BigNum) *BigNum {
	bnDestSize := MaxInt(bn.size, bnt.size) + 1
	bn.prefixZeroPadding(bnDestSize)
	bnt.prefixZeroPadding(bnDestSize)
	if signRes := bn.compareSign(bnt); signRes != eqByte {
		return bn.one2OneOpt('+', bnt)
	} else {
		numRes := bn.compareNum(bnt)
		if numRes == eqByte {
			bn.setZero()
			return bn
		} else if numRes == gtByte {
			return bn.one2OneOpt('-', bnt)
		} else {
			// whether to change bnt
			//bntCopy := bnt.makeCopy()
			bnt.one2OneOpt('-', bn)
			bn.setSelf(bnt)
			bn.resetSign()
			//bnt.setSelf(bntCopy)
			return bn
		}
	}
}

func (bn *BigNum) Mul(bnt *BigNum) *BigNum {
	return bn.simpleMul(bnt)
}

// Test Error *
func (bn *BigNum) simpleMul(bnt *BigNum) *BigNum {
	var (
		bnRSize  = bn.size + bnt.size
		bnRSign  byte
		bnResult = &BigNum{}
	)
	if bn.sign == bnt.sign {
		bnRSign = signBytePositive
	} else {
		bnRSign = signByteNegative
	}
	bnResult.setValues(make([]byte, bnRSize), bnRSize, bnRSign, baseByteDec)
	for idx1 := bn.size - 1; idx1 >= 0; idx1-- {
		for idx2 := bnt.size - 1; idx2 >= 0; idx2-- {
			bnResult.setElemOpt(idx1+idx2, '*', bn.getElem(idx1), bnt.getElem(idx2))
		}
	}
	return bnResult
}

func (bn *BigNum) fftMul(bnt *BigNum) *BigNum {
	return bn
}

func (bn *BigNum) Div(bnt *BigNum) *BigNum {
	return bn
}

func (bn *BigNum) effect() *BigNum {
	if !bn.isZero() {
		destString := TruncateStringPrefixZero(string(bn.slice))
		bn.setSlice([]byte(destString))
		bn.setSize(len(destString))
	}
	return bn
}

func (bn *BigNum) Display(displaySign, displayBase bool) *BigNum {
	bn.effect()
	if !bn.isZero() &&
		(displaySign ||
			bn.sign == signByteNegative) {
		fmt.Printf("%c", bn.sign)
	}
	if displayBase {
		fmt.Printf("0%c", bn.base)
	}
	fmt.Printf("%s\n", string(bn.slice))
	return bn
}
