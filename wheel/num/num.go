package num

// num stand for numeric calculation

import (
	"bfw/wheel/lang"
	"errors"
	"math"
)

func GCD(a, b int) int {
	return gcd(a, b)
}

func gcd(a, b int) int {
	a, b = max(a, b), min(a, b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return lcm(a, b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// QInverseSqrt
// stand for 1 / √x
// quick version of lang.InverseOfSqrt
func QInverseSqrt(x float32) float32 {
	const (
		half       float32 = 0.5
		threeHalfs float32 = 1.5
	)
	var (
		xBit     int32
		xHalf, y float32
	)
	xHalf = x * half
	y = x
	xBit = lang.GetBitByFloat32(y)
	//fmt.Printf("get x 32 bit, xBit:%x\n", xBit)
	//fmt.Printf("xBit right shift 1:%x\n", xBit>>1)
	xBit = 0x5f3759df - (xBit >> 1)
	//fmt.Printf("after WTF cal:%x\n", xBit)
	y = lang.GetFloat32ByBit(xBit)
	y = y * (threeHalfs - xHalf*y*y) // iteration 1st
	// if accuracy not enough
	//y = y * (threeHalfs -  xHalf*y*y) // iteration 2nd
	return y
}

func int2IntOpt(val int, opt rune, lambda int) int {
	switch opt {
	case '+':
		{
			return val + lambda
		}
	case '-':
		{
			return val - lambda
		}
	case '*':
		{
			return val * lambda
		}
	case '/':
		{
			return val / lambda
		}
	case '%':
		{
			return val % lambda
		}
	case '^':
		{
			return QPower(val, lambda)
		}
	default:
		{
			panic(errors.New("unsupported operator"))
		}
	}
}

func GetIntOpt(val int, opt rune, lambda int) int {
	return int2IntOpt(val, opt, lambda)
}

func SetIntOpt(val int, opt rune, lambda int) int {
	return int2IntOpt(val, opt, lambda)
}

// QPower
// quick int power by bit calculate
func QPower(base, exp int) int {
	res := 1
	for exp > 0 {
		if exp&1 != 0 {
			res *= base
		}
		base *= base
		exp >>= 1
	}
	return res
}

// SolveQuadraticEquationOfOneVariable
// delta = b^2-4a*c
func SolveQuadraticEquationOfOneVariable(a, b, c float64) (complex128, complex128, bool) {
	if lang.EqualFloat64Zero(a) {
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

// SolveCubicEquationOfOneVariableByCommon
// unfinished formula
func SolveCubicEquationOfOneVariableByCommon(a, b, c, d float64) (complex128, complex128, complex128, bool) {
	if lang.EqualFloat64Zero(a) {
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
	if lang.EqualFloat64Zero(delta) {
		if lang.EqualFloat64ByAccuracy(deltaPart1, -deltaPart2) {
			if lang.EqualFloat64Zero(deltaPart1) {
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
	if lang.EqualFloat64Zero(B) &&
		lang.EqualFloat64Zero(A) {
		return complexSolutionPart1, complexSolutionPart1, complexSolutionPart1, true
	}
	if lang.EqualFloat64Zero(Delta) {
		if !lang.EqualFloat64Zero(A) {
			k := B * 1.0 / A
			resPart1, resPart2 := -b*1.0/a, -k*1.0/2
			return complex(resPart1+k, 0.0), complex(resPart2, 0.0), complex(resPart2, 0.0), true
		}
	} else if Delta > 0.0 {
		tPart1, tPart2 := A*b*1.0+(-3.0*a*B)/2, 3.0*a*math.Sqrt(math.Pow(B*1.0, 2.0)-4.0*A*C)/2
		y1, y2 := tPart1+tPart2, tPart1-tPart2
		resPart1 := solutionPart1
		restPart2t1 := lang.Float64PowerPositivePower(y1, 1.0/3)
		restPart2t2 := lang.Float64PowerPositivePower(y2, 1.0/3)
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
