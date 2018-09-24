package luhn

import (
	"strconv"
	"strings"
)

//Valid returns true for valid Luhn compliant strings
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if !isValidPrelimChecks(s) {
		return false
	}
	var sum int
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		v := s[i] - '0'
		t := int(v)
		if IsSecond(l, i) {
			t = luhnDoubling(t)
		}
		sum += t
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func isValidPrelimChecks(s string) bool {
	return len(s) >= 1 && s != "0" && isDigit(s)
}

//luhnDoubling doubles the number and subtracts 9 if product > 9
func luhnDoubling(n int) int {
	p := 2 * n
	if p > 9 {
		p = p - 9
	}
	return p
}

//IsSecond checks if the rune is the next-secondth digit starting at 0
func IsSecond(l, n int) bool {
	if l%2 == 0 {
		return n == 0 || n%2 == 0
	}
	return n == 1 || n%2 == 1
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
