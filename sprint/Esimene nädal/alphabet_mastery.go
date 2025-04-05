package sprint

func AlphabetMastery(n int) string {

	if n < 0 || n > 26 {
		return "sisend vale" // return lõpetaks programmi töö. tükk aega jamasin
	}

	// Create a rune slice to hold the letters
	runes := make([]rune, n)

	// Generate runes for the first n letters of the alphabet
	for i := 0; i < n; i++ {
		runes[i] = 'a' + rune(i)
	}

	// Convert runes back to string and return
	return string(runes) //, nil
	//var result strings.Builder // strings.Builder is a built-in Go type specifically designed for efficient string concatenation.
	//for i := 0; i < n; i++ {
	//	letter := string('a' + i)  // Use the loop counter 'i'
	//	result.WriteString(letter) // WriteString() is a built-in method of the strings.Builder
	//}
	//return result.String()

}
