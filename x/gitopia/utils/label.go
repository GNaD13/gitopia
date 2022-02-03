package utils

import "fmt"

func LabelIdExists(l []uint64, val uint64) (int, bool) {
	for i, v := range l {
		if v == val {
			return i, true
		}
	}
	return 0, false
}

func JoinLabels(labels []string) string {
	res := ""
	len := len(labels)
	for i, l := range labels {
		res += fmt.Sprintf(" **%s**", l)
		if i == len-2 && len > 1 {
			res += " and"
		}
	}
	return res
}
