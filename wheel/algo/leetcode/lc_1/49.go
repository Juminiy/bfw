package lc_1

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	cnt := make([]int, 26)
	for i := 0; i < len(a); i++ {
		cnt[a[i]-'a']++
		cnt[b[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 {
			return false
		}
	}
	return true
}

func hash1(a string) int {
	h := 0
	for _, e := range a {
		h += int(e-96) * 101
	}
	return h
}

func anagram(str string) [26]int {
	cnt := [26]int{}
	for i := 0; i < len(str); i++ {
		cnt[str[i]-'a']++
	}
	return cnt
}

func groupAnagrams(strs []string) [][]string {
	g := make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		h := anagram(strs[i])
		if len(g[h]) == 0 {
			g[h] = make([]string, 0)
		}
		g[h] = append(g[h], strs[i])
	}
	a := make([][]string, 0)
	for _, v := range g {
		a = append(a, v)
	}
	return a
}
