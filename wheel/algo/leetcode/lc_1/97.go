package lc_1

// s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	return true
}
