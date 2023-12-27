package leetcode

func copyRandomList(head *Node) *Node {
	it, itCopy, id := makeNIter(head), makeNIter(&Node{}), 0
	itMap, itCopyMap, idId := make(map[*Node]int), make(map[int]*Node), make(map[int]int)
	for !it.nil() {
		cur := it.tail()
		itCopy.appendCopy(cur)
		itMap[cur] = id
		itCopyMap[id] = itCopy.tail()
		it.next()
		id++
	}
	itMap[nil] = -1
	it.reset()
	for !it.nil() {
		cur := it.tail()
		idId[itMap[cur]] = itMap[cur.Random]
		it.next()
	}
	itCopy.reset()
	itCopy.next()
	id = 0
	for !itCopy.nil() {
		cur := itCopy.tail()
		cur.Random = itCopyMap[idId[id]]
		id++
		itCopy.next()
	}
	return itCopy.head()
}
