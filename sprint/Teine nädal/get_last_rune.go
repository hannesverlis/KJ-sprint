package sprint

//peatükk 9., String manipulation ül 2


func GetLastRune(s string) rune {
    if len(s) == 0 {
        return 0 // Rune 0 for empty string
    }

    // Start from the last byte and move backward
    for i := len(s) - 1; i >= 0; i-- {
        // Check if the current byte is a continuation byte of a UTF-8 encoded rune
        if s[i]&0x80 == 0 {
            return []rune(s[i:])[0] // Convert to runes and return the first (which is the last in original)
        }
    }

    // If the loop completes, the string likely contains only single-byte characters.
    // Return the last byte as a rune
    return rune(s[len(s)-1])
}