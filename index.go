package piscine

func Index(s string, toFind string) int {
	l := len(s) - len(toFind)
	for i := 0; i <= l; i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
