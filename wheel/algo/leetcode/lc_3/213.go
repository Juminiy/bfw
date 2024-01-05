package lc_3

func robQ2(nums []int) int {
	if nl := len(nums); nl == 1 {
		return nums[0]
	}
	return max(robQ1(nums[1:]), robQ1(nums[:len(nums)-1]))
}
