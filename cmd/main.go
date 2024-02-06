package main

// -gcflags=-m
func main() {
	m := make(map[int]int)
	for i := 0; i < 1<<16; i++ {
		m[i] = i >> 1
	}
	clear(m)
}
