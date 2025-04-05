package sprint

func ReverseAlphabetValue(ch rune) rune {
	offset := ch - 'a'
	reverseOffset := 25 - offset
	reversedRune := reverseOffset + 'a'
	return reversedRune
}
