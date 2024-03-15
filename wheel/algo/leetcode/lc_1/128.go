package lc_1

// 100,4,200,1,3,2
// m[100] = []int{100}
// m[4] = []int{4}
// m[200] = []int{200}
// ...
func longestConsecutive(nums []int) int {
	maxl, m := 0, make(map[int]int)
	for _, num := range nums {
		m[num] = 1
	}
	for key, arr := range m {
		curKey := key + 1
		for arr0, e := m[curKey]; e; arr0, e = m[curKey] {
			delete(m, curKey)
			curKey++
			arr += arr0
		}
		m[key] = arr
		maxl = max(maxl, arr)
	}
	return maxl
}
