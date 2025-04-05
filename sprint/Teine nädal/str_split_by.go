package sprint
//Esitatud 6 juuni kl 10:57

func StrSplitBy(s, sep string) []string {
	var result []string //var result []string
    current := ""

    // Convert strings to runes (Unicode code points) for proper character handling
    runesS := []rune(s)
    runesSep := []rune(sep)

    if s == "" {    //Steni soovitatud koodijupp. Peale seda sain tehtud.
        return []string{}
    }
    // Iterate through runes in the string
    for i := 0; i < len(runesS); i++ {
        // Check if current position matches the separator
        if isSeparator(runesS, runesSep, i) {
            // Append the current substring (if not empty)
            if current != "" {
                result = append(result, current)
            }  
            // Reset the current substring for the next one
            current = ""
            // Move index to after the separator  // mida teeb ?
            i += len(runesSep) - 1
        } else {
            // Append the current rune to the current substring
           current += string(runesS[i])  // 
        }
    }
    // Append the last substring (if not empty)
    if current != "" {
        result = append(result, current)
    }
    return result
}

// Helper function to check if a substring matches the separator at a given position
func isSeparator(s, sep []rune, pos int) bool {
    for i := 0; i < len(sep); i++ {
        if pos+i >= len(s) || s[pos+i] != sep[i] {
            return false
        }
    }
    return true
}


// Steni kommentaar:  rakett — Today at 11:00 AM
//jah sest see jõuba sama asjani lõpuks. kiu sa result := string{} oleksid pannud siis poleks seda if statementi vaja
//mõlemal juhul saab tühja stringi puhul vastuseks []string{}