package main

import (
	"bfw/wheel/algo/luogu"
	"fmt"
)

// 9 10
// 0 0 0 0 0 0 1 0 0 0
// 0 0 0 0 0 0 0 0 1 0
// 0 0 0 1 0 0 0 0 0 0
// 0 0 1 0 0 0 0 0 0 0
// 0 0 0 0 0 0 1 0 0 0
// 0 0 0 0 0 1 0 0 0 0
// 0 0 0 1 1 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0 0
// 1 0 0 0 0 0 0 0 1 0
// 7 2 2 7 S
func main() {
	var (
		n, m, t        int
		graph          [][]bool
		ix, iy, ax, ay int
		iface          byte
	)
	fmt.Scan(&n, &m)
	graph = make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&t)
			graph[i][j] = t == 1
		}
	}
	fmt.Scan(&ix, &iy, &ax, &ay, &iface)
	iface = 'S'
	r := luogu.MakeRMI(n, m, graph, ix, iy, ax, ay, iface)
	fmt.Println("\n", r.BFS())
}
