package reverse

//String reverses a string
func String(s string) string {
	var rev string
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		rev += string(runes[i])
	}
	return rev
}
