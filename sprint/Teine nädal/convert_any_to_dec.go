package sprint

//esitatud 5 juuni kell 19:08
import (
    "math"
    "strings"
)

func ConvertAnyToDec(s string, base string) int {
    // Check if the base is valid
    if len(base) < 2 || strings.ContainsAny(base, "+-") || len(base) != len(strings.Split(base, "")) {
        return 0
    }

    // Create a lookup table for digit positions
    digitValues := make(map[rune]int)
    for i, r := range base {
        digitValues[r] = i
    }

    result := 0
    power := 0
    // Iterate through the string in reverse
    for i := len(s) - 1; i >= 0; i-- {
        r := rune(s[i])

        // Check if the digit is valid in the base
        value, ok := digitValues[r]
        if !ok {
            return 0 // Invalid digit
        }

        // Calculate the decimal contribution of the digit
        result += value * int(math.Pow(float64(len(base)), float64(power)))
        power++
    }

    return result
}