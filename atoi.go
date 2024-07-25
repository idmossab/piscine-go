package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		// condition 1:
		if s[i] < '0' || s[i] > '9' {
			// condition 2:
			if i == 0 && (s[0] == '-' || s[0] == '+') {
				if s[0] == '-' {
					sign *= -1
				}
				continue
			}
			return 0
		}
		dig := int(rune(s[i] - '0'))
		result = result*10*sign + dig
	}

	return result
}