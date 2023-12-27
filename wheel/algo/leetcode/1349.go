package leetcode

func seatsConverter(seats [][]string) [][]byte {
	st := make([][]byte, len(seats))
	for i, sl := range seats {
		st[i] = make([]byte, len(sl))
		for j, sa := range sl {
			st[i][j] = sa[0]
		}
	}
	return st
}

func maxStudents(seats [][]byte) int {
	// if seat
	return 0
}
