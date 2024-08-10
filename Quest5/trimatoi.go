package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && res == 0 {
			sign = -1
		} else if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
		}
	}
	return res * sign
}
