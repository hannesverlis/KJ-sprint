package sprint

// StrToInt (without strconv)
func StrToInt(s string) int {
    result := 0
    sign := 1
    start := 0

    // Check for negative sign
    if s[0] == '-' {
        sign = -1
        start = 1
    }

    for i := start; i < len(s); i++ {
        // Check for invalid characters (non-digits)
        if s[i] < '0' || s[i] > '9' {
            return 0 // Return 0 for invalid strings
        }
        result = result*10 + int(s[i]-'0')
    }

    return result * sign
}

// BulkAtoi (unchanged)
func BulkAtoi(arr []string) []int {
    result := make([]int, len(arr))
    for i, str := range arr {
        result[i] = StrToInt(str)
    }
    return result
}
