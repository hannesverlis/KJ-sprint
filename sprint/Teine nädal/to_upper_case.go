package sprint

//import "strings"  strings not allowed
//lõpetatud 5 juuni,  kl 15:57

func ToUpperCase(s string) string {
    var result []rune // Use a rune slice to hold converted characters
    for _, char := range s { // Loop through runes directly
        if char >= 'a' && char <= 'z' {
            char -= 32 
        }
        result = append(result, char) // Append the converted rune
    }
    return string(result) // Convert rune slice back to a string
}
