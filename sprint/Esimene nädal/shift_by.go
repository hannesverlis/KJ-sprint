package sprint

func ShiftBy(r rune, step int) rune {
	baseOffset := int(r - 'a')
	newOffset := baseOffset + step
	newOffset = (newOffset%26 + 26) % 26
	shiftedRune := rune(newOffset + int('a'))
	return shiftedRune
}
