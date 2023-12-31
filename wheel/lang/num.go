package lang

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
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

// GetRandomIntValue2 -> [1,n)
func GetRandomIntValue2(n int) int {
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
	sliceMap := make([]map[int]bool, 0)
	for bin := 0; bin < (1 << n); bin++ {
		if CountBits(bin) == k {
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

func CountBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func Odd(n int) bool {
	return n%2 == 1
}

func MinusOnePower(n int) int {
	if Odd(n) {
		return -1
	}
	return 1
}

func IsStringValueIntZero(valStr string) bool {
	strIdx, strLen := 0, len(valStr)
	for strIdx < strLen {
		if valStr[strIdx] == ' ' ||
			valStr[strIdx] == '+' ||
			valStr[strIdx] == '-' ||
			valStr[strIdx] == '.' ||
			valStr[strIdx] == '0' {
			strIdx++
		} else {
			break
		}
	}
	return strIdx == strLen
}

func MaxInt(a ...int) int {
	if len(a) > 0 {
		maxInt := a[0]
		for _, val := range a {
			if val > maxInt {
				maxInt = val
			}
		}
		return maxInt
	}
	return 0
}

func MinInt(a ...int) int {
	if len(a) > 0 {
		minInt := a[0]
		for _, val := range a {
			if val < minInt {
				minInt = val
			}
		}
		return minInt
	}
	return 0
}

func Square(x float64) float64 {
	return x * x
}

func Float64ToString(a float64, precision ...int) string {
	precisionBit := defaultFloat64Precision
	if len(precision) > 0 {
		precisionBit = precision[0]
	}
	return strconv.FormatFloat(a, byte('f'), precisionBit, 64)
}

// delta = b^2-4a*c
func SolveQuadraticEquationOfOneVariable(a, b, c float64) (complex128, complex128, bool) {
	if EqualFloat64Zero(a) {
		return 0, 0, false
	}
	delta := b*b - 4.0*a*c
	if delta >= 0 {
		sqrtDelta := math.Sqrt(delta)
		return complex((sqrtDelta-b)/(2.0*a), 0.0), complex((-sqrtDelta-b)/(2.0*a), 0.0), true
	} else {
		sqrtDeltaI := math.Sqrt(-delta)
		return complex(-b/(2.0*a), sqrtDeltaI/(2.0*a)), complex(-b/(2.0*a), -sqrtDeltaI/(2.0*a)), true
	}
}

func SolveCubicEquationOfOneVariableByCommon(a, b, c, d float64) (complex128, complex128, complex128, bool) {
	if EqualFloat64Zero(a) {
		return 0, 0, 0, false
	}
	var (
		deltaPart1           = math.Pow((b*c)/(6.0*math.Pow(a*1.0, 2))-math.Pow((b*1.0)/(3.0*a), 3)-(d*1.0)/(2.0*a), 2) // real
		deltaPart2           = math.Pow((c*1.0)/(3.0*a)-math.Pow((b*1.0)/(3.0*a), 2), 3)                                // real
		delta                = deltaPart1 + deltaPart2                                                                  // real
		sqrtDelta            = math.Sqrt(delta)                                                                         // real
		solutionPart1        = -(b * 1.0) / (3.0 * a)                                                                   // real
		solutionPart2        = math.Pow(deltaPart1+sqrtDelta, 1.0/3)                                                    // real
		solutionPart3        = math.Pow(deltaPart1-sqrtDelta, 1.0/3)                                                    // real
		complexConst1        = complex(-0.5, math.Sqrt(3.0)/2)                                                          // complex
		complexConst2        = complex(-0.5, -math.Sqrt(3.0)/2)                                                         // complex
		complexSolutionPart1 = complex(solutionPart1, 0.0)                                                              // complex
		complexSolutionPart2 = complex(solutionPart2, 0.0)                                                              // complex
		complexSolutionPart3 = complex(solutionPart3, 0.0)                                                              // complex
	)
	if EqualFloat64Zero(delta) {
		if EqualFloat64ByAccuracy(deltaPart1, -deltaPart2) {
			if EqualFloat64Zero(deltaPart1) {
				// three same real solutions
				tripleSolution := complexSolutionPart1
				return tripleSolution, tripleSolution, tripleSolution, true
			} else {
				// a real solution and two same real solutions
				singleSolution := complexSolutionPart1 + 2*complexSolutionPart2
				doubleSolution := complexSolutionPart1 - complexSolutionPart2
				return singleSolution, doubleSolution, doubleSolution, true
			}
		}
	} else if delta > 0 {
		// a real solution and a pair of conjugate complex solutions
		solve1 := complexSolutionPart1 + complexSolutionPart2 + complexSolutionPart3                             // complex
		solve2 := complexSolutionPart1 + complexConst1*complexSolutionPart2 + complexConst2*complexSolutionPart3 // complex
		solve3 := complexSolutionPart1 + complexConst2*complexSolutionPart2 + complexConst1*complexSolutionPart3 // complex
		return solve1, solve2, solve3, true
	} else if delta < 0 {
		// three different real solutions

		// wrong solution
		solve1 := complexSolutionPart1 + complexSolutionPart2 + complexSolutionPart3                             // complex
		solve2 := complexSolutionPart1 + complexConst1*complexSolutionPart2 + complexConst2*complexSolutionPart3 // complex
		solve3 := complexSolutionPart1 + complexConst2*complexSolutionPart2 + complexConst1*complexSolutionPart3 // complex
		return solve1, solve2, solve3, true
	}
	return 0, 0, 0, false
}

func SolveCubicEquationOfOneVariableBySJ(a, b, c, d float64) (complex128, complex128, complex128, bool) {
	var (
		A                    = math.Pow(b*1.0, 2.0) - 3.0*a*c
		B                    = b*c - 9.0*a*d
		C                    = math.Pow(c*1.0, 2.0) - 3.0*b*d
		Delta                = math.Pow(B*1.0, 2.0) - 4.0*A*C
		solutionPart1        = -b / (3.0 * a)
		complexSolutionPart1 = complex(solutionPart1, 0.0)
	)
	if EqualFloat64Zero(B) &&
		EqualFloat64Zero(A) {
		return complexSolutionPart1, complexSolutionPart1, complexSolutionPart1, true
	}
	if EqualFloat64Zero(Delta) {
		if !EqualFloat64Zero(A) {
			k := B * 1.0 / A
			resPart1, resPart2 := -b*1.0/a, -k*1.0/2
			return complex(resPart1+k, 0.0), complex(resPart2, 0.0), complex(resPart2, 0.0), true
		}
	} else if Delta > 0.0 {
		tPart1, tPart2 := A*b*1.0+(-3.0*a*B)/2, 3.0*a*math.Sqrt(math.Pow(B*1.0, 2.0)-4.0*A*C)/2
		y1, y2 := tPart1+tPart2, tPart1-tPart2
		resPart1 := solutionPart1
		restPart2t1 := Float64PowerPositivePower(y1, 1.0/3)
		restPart2t2 := Float64PowerPositivePower(y2, 1.0/3)
		resPart2 := (restPart2t1 + restPart2t2) / (6.0 * a)
		resPart3 := (math.Sqrt(3.0) * (restPart2t1 - restPart2t2)) / (6.0 * a)
		res1 := complex(resPart1-2.0*resPart2, 0.0)
		res2 := complex(resPart1+resPart2, resPart3)
		res3 := complex(resPart1+resPart2, -resPart3)
		return res1, res2, res3, true
	} else {
		if A > 0 {
			t := (2.0*A*b - 3.0*a*B) / (2.0 * A * math.Sqrt(A))
			if t >= -1.0 && t <= 1.0 {
				theta := math.Acos(t)
				resPart1 := solutionPart1
				resPart2 := (math.Sqrt(A) * math.Cos(theta*1.0/3)) / (3.0 * a)
				resPart3 := (math.Sqrt(3.0*A) * math.Sin(theta*1.0/3)) / (3.0 * a)
				res1 := complex(resPart1+resPart2+resPart3, 0.0)
				res2 := complex(resPart1+resPart2-resPart3, 0.0)
				res3 := complex(resPart1-2*resPart2, 0.0)
				return res1, res2, res3, true
			}
		}
	}
	return 0, 0, 0, false
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

func Float64Mod(a float64, m int) float64 {
	return float64(int(a) % m)
}

func AbsInt(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func CeilBin(a int) int {
	binCnt := int(math.Ceil(math.Log2(float64(a))))
	if binCnt >= 64 {
		panic(int64BitOutOfBoundError)
	}
	return 1 << binCnt
}

func MaxIntCeilBin(a ...int) int {
	destLen := 0
	if aLen := len(a); aLen > 0 {
		for aIdx := 0; aIdx < aLen; aIdx++ {
			destLen = MaxInt(destLen, a[aIdx])
		}
	}
	return CeilBin(destLen)
}
