package lang

import (
	"errors"
	"math"
	"math/bits"
	"math/rand"
	"unsafe"
)

const (
	float64Zero             float64 = 0.0
	float32Zero             float32 = 0.0
	defaultFloat32Accuracy          = 1e-5
	defaultFloat64Accuracy          = 1e-7
	defaultFloat32Precision int     = 5
	defaultFloat64Precision int     = 5
)

var (
	int64BitOutOfBoundError = errors.New("int64 bit is out of bound")
)

func EqualFloat32Zero(a float32) bool {
	return EqualFloat32ByAccuracy(a, float32Zero)
}

func EqualFloat64Zero(a float64) bool {
	return EqualFloat64ByAccuracy(a, float64Zero)
}

func EqualFloat32ByAccuracy(a, b float32, acc ...float32) bool {
	var rAcc float32 = 0.0
	if len(acc) == 0 {
		rAcc = defaultFloat32Accuracy
	} else {
		rAcc = acc[0]
	}
	if math.Abs(float64(a-b)) < float64(rAcc) {
		return true
	}
	return false
}

func EqualFloat64ByAccuracy(a, b float64, acc ...float64) bool {
	var rAcc float64 = 0.0
	if len(acc) == 0 {
		rAcc = defaultFloat64Accuracy
	} else {
		rAcc = acc[0]
	}
	if math.Abs(a-b) < rAcc {
		return true
	}
	return false
}

func GetTwoRandomIntValue(n int) (int, int) {
	return rand.Intn(n), rand.Intn(n)
}

// GetRandomIntValue -> [0,n)
func GetRandomIntValue(n int) int {
	return rand.Intn(n)
}

// GetRandomIntValueV2 -> [1,n)
func GetRandomIntValueV2(n int) int {
	return rand.Intn(n-1) + 1
}

func GetRandomMapValue(n, k int) map[int]bool {
	if k >= n {
		return nil
	}
	genMap := make(map[int]bool)
	for idx := 0; idx < k; idx++ {
		genVal := 0
		for ; genMap[genVal]; genVal = GetRandomIntValue(n) {
		}
		genMap[genVal] = true
	}
	return genMap
}

func GetCombinationSliceMap(n, k int) []map[int]bool {
	if n < 0 || k < 0 {
		panic(errors.New("combination number C_{n}^{m}, n, m cannot be negative"))
	}
	sliceMap := make([]map[int]bool, 0)
	for bin := 0; bin < (1 << n); bin++ {
		if bits.OnesCount64(uint64(bin)) == k {
			intMap := make(map[int]bool)
			for bit := 0; bit < n; bit++ {
				if (bin & (1 << bit)) > 0 {
					intMap[bit] = true
				}
			}
			sliceMap = append(sliceMap, intMap)
		}
	}
	return sliceMap
}

func IsOdd(n int) bool {
	return n%2 == 1
}

func MinusOnePower(n int) int {
	if IsOdd(n) {
		return -1
	} else {
		return 1
	}
}

func Float64PowerPositivePower(real float64, power float64) float64 {
	if EqualFloat64Zero(real) {
		return 0.0
	} else if real > 0.0 {
		return math.Pow(real, power)
	} else {
		return -math.Pow(math.Abs(real), power)
	}
}

// GetRandFloat64ByIntRange
// [0,n)
// [0,b-a)+a -> [a,b)
func GetRandFloat64ByIntRange(a, b int) float64 {
	return float64(rand.Intn(b-a) + a)
}

func GetRandFloat64ByInt32Range(a, b int32) float64 {
	return float64(rand.Int31n(b-a) + a)
}

func GetRandFloat64ByInt64Range(a, b int64) float64 {
	return float64(rand.Int63n(b-a) + a)
}

func GetRandFloat64ByFloat64Range(a, b float64) float64 {
	return rand.Float64()*(b-a) + a
}

func GetRandFloat64ArrayByRange(size int, a, b float64) []float64 {
	f := make([]float64, size)
	for idx := 0; idx < size; idx++ {
		f[idx] = GetRandFloat64ByFloat64Range(a, b)
	}
	return f
}

func Float64Mod(a float64, m int) float64 {
	return float64(int(a) % m)
}

func AbsInt(a int) int {
	if a < 0 {
		return ^a + 1
	} else {
		return a
	}
}

func CeilBinCnt(a int) int {
	binCnt := int(math.Ceil(math.Log2(float64(a))))
	if binCnt >= 64 {
		panic(int64BitOutOfBoundError)
	}
	return binCnt
}

func CeilBin(a int) int {
	return 1 << CeilBinCnt(a)
}

func MaxIntCeilBin(a ...int) int {
	destLen := 0
	if aLen := len(a); aLen > 0 {
		for aIdx := 0; aIdx < aLen; aIdx++ {
			destLen = max(destLen, a[aIdx])
		}
	}
	return CeilBin(destLen)
}

func Get4RandFloat64(a, b float64) (float64, float64, float64, float64) {
	return GetRandFloat64ByFloat64Range(a, b),
		GetRandFloat64ByFloat64Range(a, b),
		GetRandFloat64ByFloat64Range(a, b),
		GetRandFloat64ByFloat64Range(a, b)
}

func InverseOfSqrt(x float64) float64 {
	return 1.0 / math.Sqrt(x)
}

func GetBitByFloat32(x float32) int32 {
	return *(*int32)(unsafe.Pointer(&x))
}

func GetBitByFloat64(x float64) int64 {
	return *(*int64)(unsafe.Pointer(&x))
}

func GetFloat32ByBit(bit int32) float32 {
	return *(*float32)(unsafe.Pointer(&bit))
}

func GetFloat64ByBit(bit int64) float64 {
	return *(*float64)(unsafe.Pointer(&bit))
}

func IsIntSameSign(a, b int) bool {
	return (a > 0 && b > 0) ||
		(a < 0 && b < 0) ||
		(a == 0 && b == 0)
}

// Power2MulByBitCalculation
// base * (1 << exp)
func Power2MulByBitCalculation(base, exp int) int {
	return base << exp
}

func GetOriginNum(a int, sign bool) int {
	if sign {
		a = ^a + 1
	}
	return a
}
