package runtime

type s1 struct {
	a int
	b int
	c string
	s2
	s2p *s2
}

type s2 struct {
	f int
	m s3
}

type s3 struct {
	g int
	k string
}

func runDeepCopy(s1a []s1) []s1 {
	dc := make([]s1, len(s1a))
	copy(dc, s1a)
	return dc
}

func runIntDeepCopy(a []int) []int {
	dc := make([]int, len(a))
	copy(dc, a)
	return dc
}
