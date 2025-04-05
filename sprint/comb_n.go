package sprint


func CombN(n int) []string {
    if n <= 0 {
        return []string{}
    }

    combinations := []string{}
    generateCombinations(&combinations, "", 0, n)
    return combinations
}

func generateCombinations(combinations *[]string, current string, startDigit, remainingDigits int) {
    if remainingDigits == 0 {
        *combinations = append(*combinations, current)
        return
    }

    // Optimization: Avoid generating combinations that can't be ascending
    for digit := startDigit; digit <= 9-remainingDigits+1; digit++ {
        generateCombinations(combinations, current+string(digit+'0'), digit+1, remainingDigits-1)
    }
}