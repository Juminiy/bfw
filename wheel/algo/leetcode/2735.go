package leetcode

func minCost(nums []int, x int) int64 {
	minIdx, minVal := 0, int(1e9+1)
	var cost int64 = 0

	for i, val := range nums {
		if val < minVal {
			minIdx = i
			minVal = val
		}
	}
	nums, rPart := nums[minIdx:], nums[:minIdx]
	nums = append(nums, rPart...)

	return cost
}

// 我们想以最低价购买所有巧克力，转一次消耗x，将一个转到最低位置上去，每次只能左移一次 \n
// [1 15 20], 5 \n
// 如果旋转消耗大于直接购买，我们直接购买 \n
// [1 2 3], 4 \n

// 但是现在出现了一个问题，O(n)遍历数组，计算当前买和不买状态时，就已经影响了后面的状态，
// 如下面的数据
// 如果第一个直接买，第二个直接买，第三个转到第一个位置上后，总代价是：1+2+4+4+1=12
// 如果第一个直接买，第二个直接买，第三个转一下，总代价是：1+2+4+2 =9
// 如果第一个直接买，遍历到第二个，我们全款直接拿下2和3，总代价是：1+1+4+2 = 8
// 显然我们当前的操作影响了后续操作，动态规划解决不得，而且单纯比较价格和当前最低消耗是片面的，贪心也解决不得，O(n^2)暴力先试一下数据强度
// [1 2 20] ,x = 4
