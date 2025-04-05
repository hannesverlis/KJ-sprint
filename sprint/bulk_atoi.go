package sprint




// StrToInt (Modified without Unicode)
func StrToInt(s string) int {
    result := 0
    sign := 1
    index := 0

    if len(s) == 0 {
        return 0
    }

    // Handle potential leading spaces (ASCII 32)
    for index < len(s) && s[index] == 32 { // Space character in ASCII
        index++
    }

    // Handle sign (ASCII 43 for '+', 45 for '-')
    if s[index] == 45 {
        sign = -1
        index++
    } else if s[index] == 43 {
        index++ // Positive sign; continue
    }

    // Handle potential invalid characters or overflow
    for index < len(s) {
        if s[index] < 48 || s[index] > 57 {  // Check if not in '0'-'9' range (ASCII)
            return 0 // Or handle as error; in this case, we return 0
        }

        digit := int(s[index] - 48) // Convert ASCII to numerical value
        result = result*10 + digit

        index++
    }

    return sign * result
}

// BulkAtoi: Convert Array of Strings to Integers
func BulkAtoi(arr []string) []int {
    result := make([]int, len(arr)) // Initialize result array with the same length

    for i, str := range arr {
        result[i] = StrToInt(str)
    }

    return result
}
