package sprint


func BalancedParentheses(input string) bool {
    stack := []rune{}  // Stack to store opening parentheses
    mapping := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, char := range input {
        switch char {
        case '(', '[', '{':
            stack = append(stack, char) // Push opening parenthesis
        case ')', ']', '}':
            if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
                return false // Mismatch or empty stack
            }
            stack = stack[:len(stack)-1] // Pop matching parenthesis
        }
    }

    return len(stack) == 0 // All parentheses should be matched
}