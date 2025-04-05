package sprint 



func IsPalindrome(s string) bool {
    left, right := 0, len(s)-1

    for left < right {
        // Skip non-alphanumeric characters on the left side
        for left < right && !isAlphaNumeric(s[left]) {
            left++
        }
        // Skip non-alphanumeric characters on the right side
        for left < right && !isAlphaNumeric(s[right]) {
            right--
        }

        // Case-insensitive comparison (convert to lowercase)
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }

        left++
        right--
    }

    return true
}

// Helper function to check if a byte is alphanumeric
func isAlphaNumeric(b byte) bool {
    return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// Helper function to convert a byte to lowercase
func toLower(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        return b + 32 // Convert uppercase to lowercase
    }
    return b
}