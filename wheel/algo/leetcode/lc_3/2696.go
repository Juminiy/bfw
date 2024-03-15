package lc_3

func minLength(s string) int {
	//bs := unsafe.Slice(unsafe.StringData(s), len(s))
	bs := []byte(s)
	check := func() int {
		for i := 0; i < len(bs)-1; i++ {
			if (bs[i] == 'A' && bs[i+1] == 'B') ||
				(bs[i] == 'C' && bs[i+1] == 'D') {
				return i
			}
		}
		return -1
	}
	for i := check(); i != -1; i = check() {
		bs = append(bs[:i], bs[i+2:]...)
	}
	return len(bs)
}
