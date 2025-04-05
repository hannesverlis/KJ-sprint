package sprint

func DigitalRoot(n int) int {
    for n >= 10 {
        sum := 0
        for n > 0 {
            sum += n % 10  // Extract the last digit
            n /= 10       // Remove the last digit
        }
        n = sum // Update n with the sum of digits
    }
    return n 
}     