package lc_1

import "sort"

type i2Arr [][]int

func (i i2Arr) Len() int {
	return len(i)
}
func (i i2Arr) Swap(p, q int) {
	i[p], i[q] = i[q], i[p]
}
func (i i2Arr) Less(p, q int) bool {
	if i[p][0] == i[q][0] {
		return i[p][1] < i[q][1]
	}
	return i[p][0] < i[q][0]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(i2Arr(intervals))
	m := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(m) == 0 || m[len(m)-1][1] < intervals[i][0] {
			m = append(m, intervals[i])
		} else {
			m[len(m)-1][1] = max(m[len(m)-1][1], intervals[i][1])
		}
	}
	return intervals
}

//for (int i = 0; i < intervals.size(); ++i) {
//int L = intervals[i][0], R = intervals[i][1];
//if (!merged.size() || merged.back()[1] < L) {
//merged.push_back({L, R});
//}
//else {
//merged.back()[1] = max(merged.back()[1], R);
//}
//}

func canMerge(p, q []int) ([]int, bool) {
	if q[0] <= p[1] {
		return []int{p[0], q[1]}, true
	}
	return nil, false
}
