package sprint

func FactorialIterative(n int) int {
    // Error handling: Factorial is not defined for negative numbers or overflows
    if n < 0 || n > 20 { // int overflows above 20! in Go
        return 0 
    }

    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}