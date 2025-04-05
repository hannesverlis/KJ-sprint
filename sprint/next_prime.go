package sprint

func NextPrime(n int) int {
    if n <= 1 {
        return 2 // Base case: 2 is the smallest prime
    }

    if n%2 == 0 {
        n++ // If even, increment to odd number
    }

    for {
        if isPrime(n) {
            return n
        }
        n += 2 // Check only odd numbers
    }
}

func isPrime(n int) bool {
    if n <= 3 {
        return n > 1
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }

    i := 5
    for i*i <= n {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
        i += 6
    }
    return true
}
