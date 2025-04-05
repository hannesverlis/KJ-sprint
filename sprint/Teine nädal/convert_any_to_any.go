package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
    // Validate input bases (2 to 36)
    if !isValidBase(baseFrom) || !isValidBase(baseTo) {
        return "Invalid base"
    }

    // Convert from the original base to decimal (base 10)
    decimalValue := 0.0
    baseFromLen := len(baseFrom)
    for i, digit := range nbr {
        pos := strings.IndexRune(baseFrom, digit)
        if pos == -1 {
            return "Invalid number"
        }
        decimalValue += float64(pos) * math.Pow(float64(baseFromLen), float64(len(nbr)-1-i))
    }

    // Handle case where input is zero
    if decimalValue == 0 {
        return string(baseTo[0]) 
    }

    // Convert from decimal to the target base
    result := ""
    baseToLen := len(baseTo)
    for decimalValue > 0 {
        remainder := int(decimalValue) % baseToLen
        result = string(baseTo[remainder]) + result // Prepend the digit
        decimalValue = math.Floor(decimalValue / float64(baseToLen))
    }

    return result
}

// Helper function to check if a base is valid (2-36)
func isValidBase(base string) bool {
    baseLen := len(base)
    return baseLen >= 2 && baseLen <= 36 && len(strings.Trim(base, base)) == 0
}