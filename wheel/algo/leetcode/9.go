package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xp, xa := x, make([]int, 0)
	for xp > 0 {
		xa = append(xa, xp%10)
		xp /= 10
	}
	for _, e := range xa {
		xp = xp*10 + e
	}
	return xp == x
}
