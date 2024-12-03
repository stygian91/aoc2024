package q03

func parseNumber(str string) (string, bool) {
	idx := 0

	returnRes := func() (string, bool) {
		res := str[0:idx]
		if len(res) == 0 {
			return "", false
		}

		return res, true
	}

	for i, r := range str {
		idx = i
		if r < '0' || r > '9' {
			return returnRes()
		}
	}

	return returnRes()
}
