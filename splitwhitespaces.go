package piscine

func SplitWhiteSpaces(s string) []string {
	var table []string
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
		} else if res != "" {
			table = append(table, res)
			res = ""
		}
	}
	table = append(table, res)
	return table
}
