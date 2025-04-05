package sprint


import "fmt"

func StrCompress(input string) string {
    if len(input) <= 1 {
        return input // No compression needed
    }

    var compressed string
    count := 1
    currentChar := input[0]

    for i := 1; i < len(input); i++ {
        if input[i] == currentChar {
            count++
        } else {
            if count > 1 {
                compressed += fmt.Sprintf("%d%c", count, currentChar)
            } else {
                compressed += fmt.Sprintf("%c", currentChar) // Don't add count if it's 1
            }
            count = 1
            currentChar = input[i]
        }
    }

    // Handle the last character sequence
    if count > 1 {
        compressed += fmt.Sprintf("%d%c", count, currentChar)
    } else {
        compressed += fmt.Sprintf("%c", currentChar)
    }

    return compressed
}