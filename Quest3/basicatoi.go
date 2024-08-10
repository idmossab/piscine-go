package piscine

func BasicAtoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		dig := int(rune(s[i] - '0'))
		result = result*10 + dig
	}

	return result
}
