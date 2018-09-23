package isogram

import "strings"

//IsIsogram returns true if s is an isogram
func IsIsogram(s string) bool {
	var charMap = make(map[rune]struct{})
	s = strings.ToLower(s)
	for _, c := range s {
		if string(c) != "-" && string(c) != " " {
			if _, ok := charMap[c]; ok {
				return false
			}
			charMap[c] = struct{}{}
		}
	}
	return true
}
