package sprint

func NbrBase(n int, base string) string {
    if len(base) < 2 {
        return "NV"
    }

    uniqueChars := make(map[rune]bool)
    for _, char := range base {
        if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) { 
            return "NV"
        }
        if uniqueChars[char] { // Check for duplicates
            return "NV"
        }
        uniqueChars[char] = true
    }

    isNegative := n < 0
    if isNegative {
        n = -n
    }

    result := ""
    baseLen := len(base)
    for n > 0 {
        remainder := n % baseLen
        result = string(base[remainder]) + result
        n /= baseLen
    }

    if isNegative {
        result = "-" + result
    }

    return result
}