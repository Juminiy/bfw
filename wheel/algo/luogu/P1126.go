package luogu

import (
	"bfw/wheel/adt"
	"errors"
	"fmt"
)

var (
	rFaceError = errors.New("robot face error")
)

type RMI struct {
	graph    [][]bool
	visited  [][]bool
	ini, aim rNode
	n, m     int
}

func MakeRMI(n, m int, graph [][]bool, ix, iy, ax, ay int, iface byte) *RMI {
	r := &RMI{graph: graph, ini: rNode{x: ix, y: iy, face: iface}, aim: rNode{x: ax, y: ay}, n: n, m: m}
	return r.make()
}

func (r *RMI) BFS() int {
	res := r.bfs()
	if res.step == 0 &&
		res.prev == nil {
		return -1
	} else {
		r.aim.print()
		res.printPath()
		return res.step
	}
}

func (r *RMI) bfs() rRes {
	q := adt.GenericQueue[rNode]{}
	q.Push(r.ini)
	for !q.Empty() {
		t := q.Front()
		q.Pop()
		if t.equal(r.aim) {
			return t.rRes
		}
		tL := t.ins()
		for _, tE := range tL {
			if r.valNode(tE) {
				q.Push(tE)
			}
		}
		r.setVis(t.x, t.y)
		//r.debugQ(q)
	}
	return rRes{}
}

func (r *RMI) valNode(n rNode) bool {
	return r.valBound(n.x, n.y) &&
		r.valBlock(n.x, n.y) &&
		!r.visited[n.x][n.y]
}

// x,y is block NW vec
func (r *RMI) valBlock(x, y int) bool {
	cA, cB, cC, cD := true, true, true, true
	if r.valBound(x, y) {
		cA = !r.graph[x][y]
	}
	if r.valBound(x, y+1) {
		cB = !r.graph[x][y+1]
	}
	if r.valBound(x+1, y) {
		cC = !r.graph[x+1][y]
	}
	if r.valBound(x+1, y+1) {
		cD = !r.graph[x+1][y+1]
	}
	return cA && cB && cC && cD
}

func (r *RMI) valBound(x, y int) bool {
	return x >= 1 && x < r.n &&
		y >= 1 && y < r.m
}

func (r *RMI) make() *RMI {
	r.visited = nil
	r.visited = make([][]bool, r.n)
	for i := 0; i < r.n; i++ {
		r.visited[i] = make([]bool, r.m)
	}
	return r
}

func (r *RMI) setVis(x, y int) {
	if r.valBound(x, y) {
		r.visited[x][y] = true
	}
}

func (r *RMI) debugQ(q adt.GenericQueue[rNode]) {
	qa := q.GetSlice()
	for _, qe := range qa {
		qe.print()
		fmt.Print(" ")
	}
	fmt.Println()
}

func (r *RMI) aStar() {

}

// x,y is rNode central
type rNode struct {
	x, y int
	face byte
	rRes
}

func (r *rNode) equal(rt rNode) bool {
	return r.x == rt.x &&
		r.y == rt.y
}

func (r *rNode) ins() []rNode {
	rNodeList := []rNode{
		r.run(3),
		r.run(2),
		r.run(1),
		{r.x, r.y, left(r.face), makerRes(r.step+1, r)},
		{r.x, r.y, right(r.face), makerRes(r.step+1, r)},
	}
	return rNodeList
}

func (r *rNode) run(step int) rNode {
	switch r.face {
	case 'N':
		{
			return rNode{r.x - step, r.y, r.face, makerRes(r.step+1, r)}
		}
	case 'S':
		{
			return rNode{r.x + step, r.y, r.face, makerRes(r.step+1, r)}
		}
	case 'W':
		{
			return rNode{r.x, r.y - step, r.face, makerRes(r.step+1, r)}
		}
	case 'E':
		{
			return rNode{r.x, r.y + step, r.face, makerRes(r.step+1, r)}
		}
	default:
		{
			panic(rFaceError)
		}
	}
}

func (r *rNode) print() {
	fmt.Printf("(%d,%d,%c)", r.x, r.y, r.face)
}

type rRes struct {
	step int
	prev *rNode
}

func makerRes(step int, prev *rNode) rRes {
	return rRes{step, prev}
}

func (r *rRes) printPath() {
	prev := r.prev
	for prev != nil {
		fmt.Print("<-")
		prev.print()
		prev = prev.prev
	}
}

func left(face byte) byte {
	return turn(face, true)
}

func right(face byte) byte {
	return turn(face, false)
}

func turn(face byte, left bool) byte {
	switch face {
	case 'N':
		{
			if left {
				face = 'W'
			} else {
				face = 'E'
			}
		}
	case 'S':
		{
			if left {
				face = 'E'
			} else {
				face = 'W'
			}
		}
	case 'W':
		{
			if left {
				face = 'S'
			} else {
				face = 'N'
			}
		}
	case 'E':
		{
			if left {
				face = 'N'
			} else {
				face = 'S'
			}
		}
	default:
		{
			panic(rFaceError)
		}
	}
	return face
}
