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
	visited  [][][]bool
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
		for i, tE := range tL {
			if r.valNode(tE) {
				q.Push(tE)
			} else {
				if i >= 2 {
					break
				}
			}
		}
		r.setVis(t.x, t.y, t.face)
		//r.debugQ(q)
	}
	return rRes{}
}

func (r *RMI) valNode(n rNode) bool {
	return r.valRNodeCentral(n.x, n.y) &&
		r.valBlock(n.x, n.y) &&
		!r.visited[n.x][n.y][faceI(n.face)]
}

// x,y is central
// case 2*2 = 4
func (r *RMI) valBlock(x, y int) bool {
	for px := x - 1; px < x+1; px++ {
		for py := y - 1; py < y+1; py++ {
			if r.valBound(px, py) {
				if r.graph[px][py] {
					return false
				}
			}
		}
	}
	return true
}

func (r *RMI) valRNodeCentral(x, y int) bool {
	return x >= 1 && x < r.n &&
		y >= 1 && y < r.m
}

func (r *RMI) valBound(x, y int) bool {
	return x >= 0 && x < r.n &&
		y >= 0 && y < r.m
}

func (r *RMI) make() *RMI {
	r.visited = nil
	r.visited = make([][][]bool, r.n)
	for i := 0; i < r.n; i++ {
		r.visited[i] = make([][]bool, r.m)
		for j := 0; j < r.m; j++ {
			r.visited[i][j] = make([]bool, 4)
		}
	}
	return r
}

func (r *RMI) setVis(x, y int, face byte) {
	if r.valRNodeCentral(x, y) {
		r.visited[x][y][faceI(face)] = true
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
		{r.x, r.y, left(r.face), makerRes(r.step+1, r)},
		{r.x, r.y, right(r.face), makerRes(r.step+1, r)},
		r.run(1),
		r.run(2),
		r.run(3),
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

func faceI(face byte) int {
	switch face {
	case 'N':
		{
			return 0
		}
	case 'S':
		{
			return 1
		}
	case 'W':
		{
			return 2
		}
	case 'E':
		{
			return 3
		}
	default:
		{
			panic(rFaceError)
		}
	}
}
