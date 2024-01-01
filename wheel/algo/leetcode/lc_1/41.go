package lc_1

// 7,8,9,11,12
func firstMissingPositive(nums []int) int {
	n := len(nums)
	intAbs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}

	inBound := func(x int) bool {
		return x > 0 && x < n+1
	}
	for i, _ := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for _, num := range nums {
		num = intAbs(num)
		if inBound(num) {
			nums[num-1] = -intAbs(nums[num-1])
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}

	return n + 1
}

// 3 4 -1 1
// 3 4 5 1
// -3 4 -5 -1
