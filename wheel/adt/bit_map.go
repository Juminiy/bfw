package adt

import (
	"cmp"
	"errors"
)

// big number
// 1e9 billion
// 1e12 trillion
// 1e15 quadrillion
// 1e18 quintillion
const (
	bitMapCountZero    uint64 = 0
	bitMapCountOne     uint64 = 1
	bitMapCountInf     uint64 = 0xFFFFFFFFFFFFFFFF
	bitMapValNil       int64  = 0
	bitMapValMinAbs    int64  = 0
	bitMapValMaxAbs    int64  = 0x7FFFFFFFFFFFFFFF
	bitValBitByteStand uint64 = 3
	bitValBitI64Stand  uint64 = 6
	bitValBitZero      uint64 = 0
	bitValBitByte      uint64 = 8  // enough for maxVal 255
	bitValBitWord      uint64 = 16 // enough for maxVal 65535
	bitValBitTByte     uint64 = 24 // enough for maxVal 16777216
	bitValBitDWord     uint64 = 32 // enough for 4.2 Billion
	bitValBitTWord     uint64 = 48 // enough for 281 Trillion
	bitValBitQWord     uint64 = 64 // enough for 18 Quintillion.
	bitValZero         int64  = 0
	bitValMax          uint64 = 0xFFFFFFFFFFFFFFFF
	bitValMin          uint64 = 0x0
	bitValSignMax      int64  = 0x7FFFFFFFFFFFFFFF
	bitValSignMin      int64  = -1 << 63
)

var (
	unSupportedMinusValError = errors.New("you set signMode = false, please use other bitMap to put minus value")
	unSupportedValCntError   = errors.New("you set cntMode = false, please use other bitMap to put you value")
)

type BitMap struct {
	i64Bit            []uint64
	minusI64Bit       []uint64
	i64BitValCnt      []*bitValCnt
	minusI64BitValCnt []*bitValCnt
	maxBitValBit      uint64
	maxVal            int64
	minVal            int64
	totalSize         uint64
	distinctSize      uint64
	signMode          bool
	cntMode           bool
	dDMode            *discreteDistribute
}

type bitValCnt struct {
	// none used
	// i64Bit uint64
	// size = 64
	valCnt []uint64
}

type discreteDistribute struct {
	PosMaxAbs int64
	PosMinAbs int64
	NegMaxAbs int64
	NegMinAbs int64
}

func makeDDMode() *discreteDistribute {
	dD := &discreteDistribute{}
	dD.PosMaxAbs = bitMapValMinAbs
	dD.PosMinAbs = bitMapValMaxAbs
	dD.NegMaxAbs = bitMapValMinAbs
	dD.NegMinAbs = bitMapValMaxAbs
	return dD
}

func (dd *discreteDistribute) dd(val ...int64) {
	dd.PosMinAbs = minVar(absI64Var, dd.PosMinAbs, val...)
	dd.PosMaxAbs = maxVar(absI64Var, dd.PosMaxAbs, val...)
	dd.NegMinAbs = minVar(absI64Var, dd.NegMinAbs, val...)
	dd.NegMaxAbs = maxVar(absI64Var, dd.NegMaxAbs, val...)
}

func makeBitValCnt() *bitValCnt {
	bvc := &bitValCnt{}
	return bvc.make()
}

func (bvc *bitValCnt) make() *bitValCnt {
	bvc.valCnt = nil
	bvc.valCnt = make([]uint64, 1<<bitValBitI64Stand)
	return bvc
}
func (bvc *bitValCnt) makeCopy() *bitValCnt {
	bvcCopy := &bitValCnt{}
	bvcCopy.make()
	copy(bvcCopy.valCnt, bvc.valCnt)
	return bvcCopy
}

func MakeBitMap(signMode, cntMode bool, predictVal ...int64) *BitMap {
	bitMap := &BitMap{}
	var (
		destVal int64 = 1 << bitValBitByte
	)
	destVal = maxVar(absI64Var, destVal, predictVal...)
	bitMap.setMaxBitValBit(destVal)
	bitMap.resetMinMaxVal()
	return bitMap.make(signMode, cntMode)
}

func (bitMap *BitMap) make(signMode, eachCntMode bool) *BitMap {
	bitMap.signMode = signMode
	bitMap.cntMode = eachCntMode
	bitMap.alloc(bitMap.maxBitValBit)
	return bitMap
}

func (bitMap *BitMap) setMaxBitValBit(val ...int64) {
	var (
		destVal int64
	)
	destVal = maxVar(absI64Var, destVal, val...)
	if destVal < 1<<bitValBitByte {
		bitMap.maxBitValBit = bitValBitByte
		return
	}
	if destVal < 1<<bitValBitWord {
		bitMap.maxBitValBit = bitValBitWord
		return
	}
	if destVal < 1<<bitValBitTByte {
		bitMap.maxBitValBit = bitValBitTByte
		return
	}
	if destVal < 1<<bitValBitDWord {
		bitMap.maxBitValBit = bitValBitDWord
		return
	}
	if destVal < 1<<bitValBitTWord {
		bitMap.maxBitValBit = bitValBitTWord
		return
	}
	if destVal < bitValSignMax {
		bitMap.maxBitValBit = bitValBitQWord
		return
	} else {
		panic(errors.New("too big number"))
	}
}

// makeCopy
// cost too much
func (bitMap *BitMap) makeCopy() *BitMap {
	bitMapCopy := &BitMap{}
	bitMapCopy.maxBitValBit = bitMap.maxBitValBit

	return bitMapCopy
}

func (bitMap *BitMap) setDDMode() {
	bitMap.dDMode = makeDDMode()
}

// state
// noneMode 0b00
// signMode 0b01
// cntMode  0b10
// mixMode  0b11
func (bitMap *BitMap) state() uint8 {
	var (
		state uint8 = 0b00
	)
	if bitMap.signMode {
		state |= 0b01
	}
	if bitMap.cntMode {
		state |= 0b10
	}
	return state
}

// i64Bit[7] -> [511,510,...,450,449,448]
// ...
// ...
// ...
// i64Bit[1] -> [127,126,...,66,65,64]
// i64Bit[0] -> [63,62,...,2,1,0]
func (bitMap *BitMap) alloc(bitValBit uint64) {
	bitMap.dealloc()
	bitCnt := bitValBit - bitValBitI64Stand
	switch bitMap.state() {
	case 0b00:
		{
			bitMap.i64Bit = make([]uint64, 1<<bitCnt)
		}
	case 0b01:
		{
			bitMap.i64Bit = make([]uint64, 1<<bitCnt)
			bitMap.minusI64Bit = make([]uint64, 1<<bitCnt)
		}
	case 0b10:
		{
			bitMap.i64Bit = make([]uint64, 1<<bitCnt)
			bitMap.i64BitValCnt = make([]*bitValCnt, 1<<bitCnt)
		}
	case 0b11:
		{
			bitMap.i64Bit = make([]uint64, 1<<bitCnt)
			bitMap.minusI64Bit = make([]uint64, 1<<bitCnt)
			bitMap.i64BitValCnt = make([]*bitValCnt, 1<<bitCnt)
			bitMap.minusI64BitValCnt = make([]*bitValCnt, 1<<bitCnt)
		}
	default:
		{
			panic(errors.New("unSupported bitMap State"))
		}
	}
	bitMap.maxBitValBit = bitValBit
}

func (bitMap *BitMap) incI64Bit(bitValBit uint64, sign bool) {
	prevBitCnt := bitMap.maxBitValBit - bitValBitI64Stand
	destBitCnt := bitValBit - bitValBitI64Stand
	padBit := make([]uint64, (1<<destBitCnt)-(1<<prevBitCnt))
	if !sign {
		bitMap.i64Bit = append(bitMap.i64Bit, padBit...)
	} else {
		bitMap.minusI64Bit = append(bitMap.minusI64Bit, padBit...)
	}
}

func (bitMap *BitMap) decI64Bit(bitValBit uint64, sign bool) {
	destBitCnt := bitValBit - bitValBitI64Stand
	if !sign {
		bitMap.i64Bit = bitMap.i64Bit[:1<<destBitCnt]
	} else {
		bitMap.minusI64Bit = bitMap.minusI64Bit[:1<<destBitCnt]
	}
}

func (bitMap *BitMap) incI64BitValCnt(bitValBit uint64, sign bool) {
	prevBitCnt := bitMap.maxBitValBit - bitValBitI64Stand
	destBitCnt := bitValBit - bitValBitI64Stand
	padBit := make([]*bitValCnt, (1<<destBitCnt)-(1<<prevBitCnt))
	if !sign {
		bitMap.i64BitValCnt = append(bitMap.i64BitValCnt, padBit...)
	} else {
		bitMap.minusI64BitValCnt = append(bitMap.minusI64BitValCnt, padBit...)
	}
}

func (bitMap *BitMap) decI64BitValCnt(bitValBit uint64, sign bool) {
	destBitCnt := bitValBit - bitValBitI64Stand
	if !sign {
		bitMap.i64BitValCnt = bitMap.i64BitValCnt[:1<<destBitCnt]
	} else {
		bitMap.minusI64BitValCnt = bitMap.minusI64BitValCnt[:1<<destBitCnt]
	}
}

func (bitMap *BitMap) realloc(bitValBit uint64, truncate ...bool) {
	if bitValBit > bitMap.maxBitValBit {
		switch bitMap.state() {
		case 0b00:
			{
				bitMap.incI64Bit(bitValBit, false)
			}
		case 0b01:
			{
				bitMap.incI64Bit(bitValBit, false)
				bitMap.incI64Bit(bitValBit, true)
			}
		case 0b10:
			{
				bitMap.incI64Bit(bitValBit, false)
				bitMap.incI64BitValCnt(bitValBit, false)
			}
		case 0b11:
			{
				bitMap.incI64Bit(bitValBit, false)
				bitMap.incI64Bit(bitValBit, true)
				bitMap.incI64BitValCnt(bitValBit, false)
				bitMap.incI64BitValCnt(bitValBit, true)
			}
		default:
			{
				panic(errors.New("unSupported bitMap State"))
			}
		}
	} else if bitValBit == bitMap.maxBitValBit {

	} else {
		if len(truncate) > 0 && truncate[0] {
			switch bitMap.state() {
			case 0b00:
				{
					bitMap.decI64Bit(bitValBit, false)
				}
			case 0b01:
				{
					bitMap.decI64Bit(bitValBit, false)
					bitMap.decI64Bit(bitValBit, true)
				}
			case 0b10:
				{
					bitMap.decI64Bit(bitValBit, false)
					bitMap.decI64BitValCnt(bitValBit, false)
				}
			case 0b11:
				{
					bitMap.decI64Bit(bitValBit, false)
					bitMap.decI64Bit(bitValBit, true)
					bitMap.decI64BitValCnt(bitValBit, false)
					bitMap.decI64BitValCnt(bitValBit, true)
				}
			default:
				{
					panic(errors.New("unSupported bitMap State"))
				}
			}
		}
	}
	bitMap.maxBitValBit = bitValBit
}

func (bitMap *BitMap) dealloc() {
	bitMap.i64Bit = nil
	bitMap.minusI64Bit = nil
	// TODO: whether need to do deeply dealloc
	bitMap.i64BitValCnt = nil
	bitMap.minusI64BitValCnt = nil
}

func (bitMap *BitMap) expand() {
	switch bitMap.maxBitValBit {
	case bitValBitZero:
		{
			bitMap.realloc(bitValBitByte)
		}
	case bitValBitByte:
		{
			bitMap.realloc(bitValBitWord)
		}
	case bitValBitWord:
		{
			bitMap.realloc(bitValBitTByte)
		}
	case bitValBitTByte:
		{
			bitMap.realloc(bitValBitDWord)
		}
	case bitValBitDWord:
		{
			bitMap.realloc(bitValBitTWord)
		}
	case bitValBitTWord:
		{
			bitMap.realloc(bitValBitQWord)
		}
	case bitValBitQWord:
		{
			panic(errors.New("cannot allocate any more memory, OOM"))
		}
	default:
		{
			panic(errors.New("maxBitValBit is not regular"))
		}
	}
}

func (bitMap *BitMap) setNull() {
	bitMap.dealloc()
	bitMap.maxBitValBit = bitValBitZero
	bitMap.maxVal = bitMapValNil
	bitMap.minVal = bitMapValNil
	bitMap.totalSize = bitMapCountZero
	bitMap.distinctSize = bitMapCountZero
	bitMap.signMode = false
	bitMap.cntMode = false
}

func (bitMap *BitMap) assign(bitMapT *BitMap) {
	bitMap.i64Bit = bitMapT.i64Bit
	bitMap.minusI64Bit = bitMapT.minusI64Bit
	bitMap.i64BitValCnt = bitMapT.i64BitValCnt
	bitMap.minusI64BitValCnt = bitMapT.minusI64BitValCnt
	bitMap.maxBitValBit = bitMapT.maxBitValBit
	bitMap.maxVal = bitMapT.maxVal
	bitMap.minVal = bitMapT.minVal
	bitMap.totalSize = bitMapT.totalSize
	bitMap.distinctSize = bitMapT.distinctSize
	bitMap.signMode = bitMapT.signMode
	bitMap.cntMode = bitMapT.cntMode
}

func (bitMap *BitMap) swap(bitMapT *BitMap) {
	bitMapTP := &BitMap{}
	bitMapTP.assign(bitMap)
	bitMap.assign(bitMapT)
	bitMapT.assign(bitMapTP)
}

func (bitMap *BitMap) resetMinMaxVal() {
	if bitMap.signMode {
		bitMap.maxVal = bitValSignMin
		bitMap.minVal = bitValSignMax
	} else {
		bitMap.maxVal = bitValZero
		bitMap.minVal = bitValSignMax
	}
}

func (bitMap *BitMap) clear() {
	bitMap.dealloc()
	bitMap.maxBitValBit = bitValBitZero
	bitMap.resetMinMaxVal()
	bitMap.totalSize = bitMapCountZero
	bitMap.distinctSize = bitMapCountZero
}

func (bitMap *BitMap) empty() bool {
	return bitMap.totalSize == bitMapCountZero
}

func (bitMap *BitMap) setMinVal(val ...int64) {
	bitMap.minVal = minVar(nilI64Var, bitMap.minVal, val...)
}

func (bitMap *BitMap) getMinVal(val ...int64) int64 {
	bitMap.setMinVal(val...)
	return bitMap.minVal
}

func (bitMap *BitMap) setMaxVal(val ...int64) {
	bitMap.maxVal = maxVar(absI64Var, bitMap.maxVal, val...)
}

func (bitMap *BitMap) getMaxVal(val ...int64) int64 {
	bitMap.setMaxVal(val...)
	return bitMap.maxVal
}

func (bitMap *BitMap) setMinMaxVal(val int64) {
	if bitMap.dDMode != nil {
		bitMap.dDMode.dd(val)
	}
	bitMap.setMinVal(val)
	bitMap.setMaxVal(val)
}

func (bitMap *BitMap) locate(x int64) (int, int64) {
	if x < 0 {
		if !bitMap.signMode {
			panic(unSupportedMinusValError)
		}
		x = ^x + 1
	}
	arrIdx := x >> bitValBitI64Stand
	bitIdx := x & 0b111111
	return int(arrIdx), bitIdx
}

// opt & 0b01 -> set
// opt & 0b10 -> rmv
func (bitMap *BitMap) optAtLocation(arrIdx int, bitIdx int64, opt uint8, signVar ...bool) (exists bool, sign bool, value int64) {
	signLen := len(signVar)
	val := int64(arrIdx<<bitValBitI64Stand) + bitIdx
	if signLen == 0 || (signLen > 0 && !signVar[0]) {
		exists = (bitMap.i64Bit[arrIdx] & (1 << bitIdx)) != 0
		if exists {
			value = val
		}
		sign = false
		if opt&0b1 != 0 {
			bitMap.i64Bit[arrIdx] |= 1 << bitIdx
		}
		if opt&0b10 != 0 {
			bitMap.i64Bit[arrIdx] &= bitValMax - 1<<bitIdx
		}
	} else if signLen > 0 && signVar[0] {
		if !bitMap.signMode {
			panic(unSupportedMinusValError)
		}
		exists = (bitMap.minusI64Bit[arrIdx] & (1 << bitIdx)) != 0
		if exists {
			value = ^val + 1
		}
		sign = true
		if opt&0b1 != 0 {
			bitMap.minusI64Bit[arrIdx] |= 1 << bitIdx
		}
		if opt&0b10 != 0 {
			bitMap.minusI64Bit[arrIdx] &= bitValMax - 1<<bitIdx
		}
	} else {
		exists = false
		sign = false
		value = bitMapValNil
	}
	return
}

// opt & 0b01 -> set -> valCnt += cnt
// opt & 0b10 -> rmv -> valCnt -= cnt
func (bitMap *BitMap) optAtLocationValCnt(arrIdx int, bitIdx int64, opt uint8, sign bool, cnt uint64) (valCnt uint64) {
	if !bitMap.cntMode {
		panic(unSupportedValCntError)
	}
	var bitObj *bitValCnt
	if !sign {
		bitObj = bitMap.i64BitValCnt[arrIdx]
	} else {
		if !bitMap.signMode {
			panic(unSupportedMinusValError)
		}
		bitObj = bitMap.minusI64BitValCnt[arrIdx]
	}
	if bitObj == nil {
		bitObj = makeBitValCnt()
	}
	valCnt = bitObj.valCnt[bitIdx]
	if opt&0b01 != 0 {
		bitObj.valCnt[bitIdx] += cnt
	}
	if opt&0b10 != 0 {
		if valCnt > cnt {
			bitObj.valCnt[bitIdx] -= cnt
		} else {
			bitObj.valCnt[bitIdx] = bitMapCountZero
		}
	}
	// write back
	if !sign {
		bitMap.i64BitValCnt[arrIdx] = bitObj
	} else {
		if !bitMap.signMode {
			panic(unSupportedMinusValError)
		}
		bitMap.minusI64BitValCnt[arrIdx] = bitObj
	}
	return
}

func (bitMap *BitMap) set(arrIdx int, bitIdx int64, signVar ...bool) {
	bitMap.optAtLocation(arrIdx, bitIdx, 0b1, signVar...)
}

func (bitMap *BitMap) rmv(arrIdx int, bitIdx int64, signVar ...bool) {
	bitMap.optAtLocation(arrIdx, bitIdx, 0b10, signVar...)
}

func (bitMap *BitMap) get(arrIdx int, bitIdx int64, signVar ...bool) (bool, uint64) {
	exists, sign, _ := bitMap.optAtLocation(arrIdx, bitIdx, 0, signVar...)
	var (
		valCnt = bitMapCountZero
	)
	if bitMap.cntMode {
		valCnt = bitMap.optAtLocationValCnt(arrIdx, bitIdx, 0, sign, 0)
	} else {
		if exists {
			valCnt = bitMapCountOne
		}
	}
	return exists, valCnt
}

func (bitMap *BitMap) getSet(arrIdx int, bitIdx int64, signVar ...bool) {
	exists, sign, _ := bitMap.optAtLocation(arrIdx, bitIdx, 0b1, signVar...)
	bitMap.totalSize++
	if !exists {
		bitMap.distinctSize++
	}
	if bitMap.cntMode {
		bitMap.optAtLocationValCnt(arrIdx, bitIdx, 0b1, sign, bitMapCountOne)
	}
}

// a bit tough
func (bitMap *BitMap) getRmv(arrIdx int, bitIdx int64, signVar ...bool) {
	exists, sign, _ := bitMap.optAtLocation(arrIdx, bitIdx, 0, signVar...)
	if !bitMap.cntMode {
		if exists {
			bitMap.rmv(arrIdx, bitIdx, sign)
			bitMap.distinctSize--
			bitMap.totalSize--
		}
	} else {
		valCnt := bitMap.optAtLocationValCnt(arrIdx, bitIdx, 0b10, sign, bitMapCountOne)
		if valCnt > 0 {
			bitMap.totalSize--
			if valCnt == 1 {
				bitMap.rmv(arrIdx, bitIdx, sign)
				bitMap.distinctSize--
			}
		}
	}
}

func (bitMap *BitMap) getVal(arrIdx int, bitIdx int64, signVar ...bool) (bool, int64) {
	exists, _, value := bitMap.optAtLocation(arrIdx, bitIdx, 0, signVar...)
	return exists, value
}

func (bitMap *BitMap) getValCnt(arrIdx int, bitIdx int64, signVar ...bool) (bool, int64, uint64) {
	exists, sign, value := bitMap.optAtLocation(arrIdx, bitIdx, 0, signVar...)
	valCnt := bitMap.optAtLocationValCnt(arrIdx, bitIdx, 0, sign, 0)
	return exists, value, valCnt
}

func (bitMap *BitMap) exists(arrIdx int, bitIdx int64, signVar ...bool) bool {
	exists, _ := bitMap.get(arrIdx, bitIdx, signVar...)
	return exists
}

func (bitMap *BitMap) attribute(x int64) (bool, int, int64) {
	sign := false
	if x < 0 {
		sign = true
	}
	arrIdx, bitIdx := bitMap.locate(x)
	// expand maxBitValBit
	if uint64(arrIdx) >= 1<<bitMap.maxBitValBit {
		bitMap.expand()
	}
	return sign, arrIdx, bitIdx
}

func (bitMap *BitMap) reduce(x int64, cnt uint64) {
	sign, arrIdx, bitIdx := bitMap.attribute(x)
	bitMap.optAtLocationValCnt(arrIdx, bitIdx, 0b10, sign, cnt)
}

func (bitMap *BitMap) batchInsert(x ...int64) {
	if xLen := len(x); xLen > 0 {
		for xIdx := 0; xIdx < xLen; xIdx++ {
			bitMap.insert(x[xIdx])
		}
	}
}

func (bitMap *BitMap) Insert(x int64) *BitMap {
	bitMap.insert(x)
	return bitMap
}

func (bitMap *BitMap) insert(x int64) {
	// attribute x
	sign, arrIdx, bitIdx := bitMap.attribute(x)

	// getSet
	bitMap.getSet(arrIdx, bitIdx, sign)

	// set maxVal, minVal
	bitMap.setMinMaxVal(x)
}

func (bitMap *BitMap) Erase(x int64) *BitMap {
	bitMap.erase(x)
	return bitMap
}

func (bitMap *BitMap) erase(x int64) {
	bitMap.reduce(x, bitMapCountInf)
}

func (bitMap *BitMap) Decline(x int64) *BitMap {
	bitMap.reduce(x, bitMapCountOne)
	return bitMap
}

func (bitMap *BitMap) Find(x int64) bool {
	return bitMap.find(x)
}

func (bitMap *BitMap) find(x int64) bool {
	sign, arrIdx, bitIdx := bitMap.attribute(x)
	return bitMap.exists(arrIdx, bitIdx, sign)
}

func (bitMap *BitMap) Count(x int64) uint64 {
	return bitMap.count(x)
}

func (bitMap *BitMap) count(x int64) uint64 {
	sign, arrIdx, bitIdx := bitMap.attribute(x)
	exists, cnt := bitMap.get(arrIdx, bitIdx, sign)
	destCnt := bitMapCountZero
	if cnt > 0 {
		destCnt = cnt
	} else {
		if exists {
			destCnt = bitMapCountOne
		}
	}
	return destCnt
}

func (bitMap *BitMap) getDistinctArray(sign bool) []int64 {
	i64Arr, i64Idx := make([]int64, bitMap.totalSize), 0
	if sign {
		maxMinusArrIdx, maxMinusBitIdx := bitMap.locate(bitMap.minVal)
		for arrIdx := maxMinusArrIdx; arrIdx >= 0; arrIdx-- {
			for bitIdx := int(maxMinusBitIdx); bitIdx >= 0; bitIdx-- {
				if exists, val := bitMap.getVal(arrIdx, int64(bitIdx), true); exists {
					i64Arr[i64Idx] = val
					i64Idx++
				}
			}
		}
	}
	maxArrIdx, maxBitIdx := bitMap.locate(bitMap.maxVal)
	for arrIdx := 0; arrIdx <= maxArrIdx; arrIdx++ {
		for bitIdx := 0; bitIdx <= int(maxBitIdx); bitIdx++ {
			if exists, val := bitMap.getVal(arrIdx, int64(bitIdx)); exists {
				i64Arr[i64Idx] = val
				i64Idx++
			}
		}
	}
	return i64Arr
}

func (bitMap *BitMap) getAllArray(sign bool) []int64 {
	i64Arr, i64Idx := make([]int64, bitMap.totalSize), 0
	if sign {
		maxMinusArrIdx, maxMinusBitIdx := bitMap.locate(bitMap.minVal)
		for arrIdx := maxMinusArrIdx; arrIdx >= 0; arrIdx-- {
			for bitIdx := int(maxMinusBitIdx); bitIdx >= 0; bitIdx-- {
				if exists, val, valCnt := bitMap.getValCnt(arrIdx, int64(bitIdx), true); exists {
					for cntI := int64(valCnt); cntI > 0; cntI-- {
						i64Arr[i64Idx] = val
						i64Idx++
					}
				}
			}
		}
	}
	var (
		minArrIdx            int
		maxArrIdx, maxBitIdx = bitMap.locate(bitMap.maxVal)
	)
	if bitMap.dDMode != nil {
		minArrIdx, _ = bitMap.locate(bitMap.dDMode.PosMinAbs)
		maxArrIdx, maxBitIdx = bitMap.locate(bitMap.dDMode.PosMaxAbs)
	}
	for arrIdx := minArrIdx; arrIdx <= maxArrIdx; arrIdx++ {
		for bitIdx := 0; bitIdx <= int(maxBitIdx); bitIdx++ {
			if exists, val, valCnt := bitMap.getValCnt(arrIdx, int64(bitIdx)); exists {
				for cntI := int64(valCnt); cntI > 0; cntI-- {
					i64Arr[i64Idx] = val
					i64Idx++
				}
			}
		}
	}
	return i64Arr
}

func (bitMap *BitMap) Array() []int64 {
	switch bitMap.state() {
	case 0b00:
		{
			return bitMap.getDistinctArray(false)
		}
	case 0b01:
		{
			return bitMap.getDistinctArray(true)
		}
	case 0b10:
		{
			return bitMap.getAllArray(false)
		}
	case 0b11:
		{
			return bitMap.getAllArray(true)
		}
	default:
		{
			panic(errors.New("unSupported bitMap State"))
		}
	}
}

func maxVar[T cmp.Ordered](f func(T) T, x T, y ...T) T {
	for idx := 0; idx < len(y); idx++ {
		x = max(x, f(y[idx]))
	}
	return x
}

func minVar[T cmp.Ordered](f func(T) T, x T, y ...T) T {
	for idx := 0; idx < len(y); idx++ {
		x = min(x, f(y[idx]))
	}
	return x
}

func nilI64Var(x int64) int64 {
	return x
}

func absI64Var(x int64) int64 {
	if x < 0 {
		x = ^x + 1
	}
	return x
}

func absUI64Var(x uint64) uint64 {
	return x
}

func absVar(x interface{}) interface{} {
	switch x.(type) {
	case int8:
		{
			i8X := x.(int8)
			if i8X < 0 {
				i8X = ^i8X + 1
			}
			return i8X
		}
	case int16:
		{
			i16X := x.(int16)
			if i16X < 0 {
				i16X = ^i16X + 1
			}
			return i16X
		}
	case int32:
		{
			i32X := x.(int32)
			if i32X < 0 {
				i32X = ^i32X + 1
			}
			return i32X
		}
	case int64:
		{
			i64X := x.(int64)
			if i64X < 0 {
				i64X = ^i64X + 1
			}
			return i64X
		}
	case int:
		{
			iX := x.(int)
			if iX < 0 {
				iX = ^iX + 1
			}
			return iX
		}
	default:
		{
			panic(errors.New("unSupported int type"))
		}
	}
	return nil
}

func BitMapSort(arr []int64, predictMaxVal int64) []int64 {
	bitMap := MakeBitMap(false, true, predictMaxVal)
	bitMap.setDDMode()
	bitMap.batchInsert(arr...)
	return bitMap.Array()
}
