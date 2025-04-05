package sprint


// ArrCountIf Function
func ArrCountIf(f func(string) bool, tab []string) int {
    count := 0
    for _, element := range tab {
        if f(element) {
            count++
        }
    }
    return count
}
func IsNumeric(s string) bool {
    for _, r := range s {
        if r < '0' || r > '9' { 
            return false 
        }
    }
    return true
}

func IsLower(s string) bool {
    for _, r := range s { 
        if r < 'a' || r > 'z' { 
            return false
        }
    }
    return true
}

func IsUpper(s string) bool {
    for _, r := range s {
        if r < 'A' || r > 'Z' {
            return false
        }
    }
    return true
}

func IsAlphanumeric(s string) bool {
    for _, r := range s {
        if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && (r < '0' || r > '9') {
            return false
        }
    }
    return true
}