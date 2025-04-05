package sprint


// FindDividend returns the first number in the range [from, to) that is divisible by divisor.
// If no such number exists, it returns -1.
func FindDividend(from, to, divisor int) int {
    for i := from; i < to; i++ {
        if i % divisor == 0 {
            return i
        }
    }
    return -1
}