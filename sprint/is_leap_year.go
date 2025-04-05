package sprint
func IsLeapYear(year int) bool {
    // Check if divisible by 4
    if year%4 == 0 {
        // Check for exceptions: centuries divisible by 100 but not 400
        if year%100 == 0 {
            return year%400 == 0
        }
        return true // Divisible by 4 but not by 100
    }
    return false // Not divisible by 4
}