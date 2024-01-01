package lc_1

// 4x+2y=t
// 2x+2y=2c
// x = t/2-c
// y = 2c-t/2
func numOfBurgers(t, c int) []int {
	x, y := t>>1-c, c<<1-t>>1
	if x >= 0 && y >= 0 &&
		t%2 == 0 {
		return []int{x, y}
	}
	return []int{}
}
