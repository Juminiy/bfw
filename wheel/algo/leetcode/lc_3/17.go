package lc_3

// 2 abc
// 3 def
// 4 ghi
// 5 jkl
// 6 mno
// 7 pqrs
// 8 tuv
// 9 wxyz
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	m := map[int][]string{
		1: {}, 2: {"a", "b", "c"}, 3: {"d", "e", "f"},
		4: {"g", "h", "i"}, 5: {"j", "k", "l"}, 6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"}, 8: {"t", "u", "v"}, 9: {"w", "x", "y", "z"},
	}
	dfs(0, "", res, digits, m)
	return res
}
func dfs(i int, s string, res []string, digits string, m map[int][]string) {
	if i == len(digits) {
		res = append(res, s)
		return
	}
	for _, si := range m[int(digits[i]-'1'+1)] {
		dfs(i+1, s+si, res, digits, m)
	}
}
