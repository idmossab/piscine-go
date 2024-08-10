package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		if table[i] > table[i+1] {
			c := table[i]
			table[i] = table[i+1]
			table[i+1] = c
			i = -1
		}
	}
}
