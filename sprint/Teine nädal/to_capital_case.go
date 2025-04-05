package sprint


func ToCapitalCase(s string) string {
    result := []byte(s)    // Convert to byte slice for efficient modification
    capitalizeNext := true // Flag to track if the next letter should be capitalized
    for i := 0; i < len(result); i++ {
        if result[i] >= 'A' && result[i] <= 'Z' || 
           result[i] >= 'a' && result[i] <= 'z' ||
           result[i] >= '0' && result[i] <= '9' {
            if capitalizeNext {
                if result[i] >= 'a' && result[i] <= 'z' {
                    result[i] -= 32 // Convert to uppercase (ASCII manipulation)
                }
                capitalizeNext = false
            } else if result[i] >= 'A' && result[i] <= 'Z' {
                result[i] += 32 // Convert to lowercase (ASCII manipulation)
            }
        } else {
            capitalizeNext = true
        }
    }
    return string(result)
}
