package sprint
//lõpetatud 5 juuni kl 18:02

func SubstrIndex(s string, toFind string) int {

    if len(toFind) == 0 {
        return 0 // Empty string is always found at index 0
    }

    for i := 0; i <= len(s)-len(toFind); i++ {
        match := true
        for j := 0; j < len(toFind); j++ {
            if s[i+j] != toFind[j] {
                match = false
                break
            }
        }
        if match {
            return i
        }
    }
    return -1 // Substring not found
}