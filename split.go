package piscine

func Split(s, sep string) []string {
	var table []string
	res := ""
	sepLen := len(sep)

	for i := 0; i < len(s); {
		//i+seplen <=len(s) bach mayfotlnach dik string li3andna
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			if res != "" {
				table = append(table, res)
				res = ""
			}
			//kan9zo b dak len dyal dik lklma liradi ndiro biha espace  bach man7sbohach manba3d
			i += sepLen
		} else {
			res += string(s[i])
			i++
		}
	}
	if res != "" {
		table = append(table, res)
	}
	return table
}
