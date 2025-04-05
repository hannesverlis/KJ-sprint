package sprint


func SplitWhitespaces(s string) []string {
	var words []string       // Slice to store the words
	var currentWord string    // Temporary buffer for building each word

	for _, r := range s { // Iterate over each rune (character) in the string
		if r == ' ' || r == '\t' || r == '\n' { // Check for whitespace
			if currentWord != "" {          // If a word was being built, add it to the slice  " " tühiku lisamine, lisas selle sõna ette
				words = append(words, currentWord)
				currentWord = ""             // Reset the buffer for the next word  // " " tühiku lisamine, lisas selle sõna ette
			}
		} else {
			currentWord += string(r)    // Append the character to the current word
		}
	}

	if currentWord != "" { // Add the last word if it wasn't followed by whitespace
		words = append(words, currentWord)
	}
	

	return words
}