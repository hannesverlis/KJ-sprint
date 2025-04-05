package sprint
//valmis 5 juuni kl 12:25
func GetFirstRune(s string) rune {
if len(s) > 0 {
	return []rune(s)[0] // Convert to runes and get the first
}
return 0 // Rune 0 represents an invalid or null rune
}