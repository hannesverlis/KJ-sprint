package sprint

//peatükk 9., String manipulation ül 3. tehtud 5 juuni 13:07

func NRune(s string, i int) rune {
    // Ensure the index is within valid bounds (0 <= i < len(s))
    if i < 0 || i >= len(s) {
        return 0 // Rune 0 for invalid index
    }

    return []rune(s)[i] // Convert to runes and access directly
}
