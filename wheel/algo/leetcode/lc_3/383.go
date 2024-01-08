package lc_3

func canConstruct(r, m string) bool {
	var cnt [26]int
	for _, me := range m {
		cnt[me-'a']++
	}
	for _, re := range r {
		cnt[re-'a']--
		if cnt[re-'a'] < 0 {
			return false
		}
	}
	return true
}
