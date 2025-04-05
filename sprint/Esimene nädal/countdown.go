package sprint

func Countdown(n int) string {
	if n < 0 || n > 9 {
		return "vale sisend"
	}

	result := ""
	for i := n; i >= 0; i -= 2 {
		if n%2 == 1 { // paaris, paaritu kontroll
			result += string('0' + i) // Convert digit to character directly
			if i > 0 {
				result += ", "
			}
			if i < 1 {
				result += "0!"
			}

		} else {
			if n%2 == 0 { // paaris, paaritu kontroll
				result += string('0' + i) // Convert digit to character directly
				if i > 0 {
					result += ", "
				}

			}

		}

		result += "!"

	}
	return result
}
