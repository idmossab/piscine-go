package piscine

func BasicAtoi2(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		// condition 1:
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		dig := int(rune(s[i] - '0'))
		result = result*10 + dig
	}

	return result
}
