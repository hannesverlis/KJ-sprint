package sprint

func FromRoman(s string) int {
    // Map to store Roman numeral values
    romanValues := map[byte]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
    }

    result := 0
    prevValue := 0 // Keep track of the previous numeral's value

    // Iterate through the Roman numeral string in reverse
    for i := len(s) - 1; i >= 0; i-- {
        currentValue := romanValues[s[i]]

        if currentValue < prevValue { 
            result -= currentValue // Subtract if the current numeral is smaller than the previous
        } else {
            result += currentValue // Add if the current numeral is greater or equal
        }

        prevValue = currentValue // Update the previous value for the next iteration
    }

    return result
}