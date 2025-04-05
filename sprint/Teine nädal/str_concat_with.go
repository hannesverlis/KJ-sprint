package sprint

func StrConcatWith(strs []string, sep string) string { // sep - separaator, koolun, koma, jne..

    if len(strs) == 0 {
        return "" // Handle an empty slice
    }

    result := strs[0] // Start with the first string
    for _, str := range strs[1:] { // Loop through the remaining strings
        result += sep + str // Append separator and string
    }
    return result
}