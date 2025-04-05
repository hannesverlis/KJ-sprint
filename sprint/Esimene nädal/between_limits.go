package sprint

func BetweenLimits(from, to rune) string {
	if from > to {
		from, to = to, from
	}

	var result []rune
	for r := from + 1; r < to; r++ {
		if r > 'z' {
			r = 'a' + (r - 'z' - 1)
		}

		if r <= to {
			result = append(result, r)
		}
	}
	return string(result)
}
