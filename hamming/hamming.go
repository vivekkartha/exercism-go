package hamming

import u "unicode/utf8"
import "errors"

//Distance returns the hamming distance between the 2 string params.
func Distance(a, b string) (i int, er error) {
	var dist = 0
	if u.RuneCountInString(a) != u.RuneCountInString(b) {
		error := errors.New("Unequal strings")
		return 0, error
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
