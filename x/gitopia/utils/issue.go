package utils

func IssueCommentExists(c []uint64, val uint64) (int, bool) {
	for i, v := range c {
		if v == val {
			return i, true
		}
	}
	return 0, false
}
