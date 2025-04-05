package main

import "fmt"

func FactorialIterative(n int) int {
    // Handle non-positive input and potential overflow
    if n < 0 {
        return 0 // Factorial isn't defined for negative numbers
    } else if n > 12 { 
        return 0 // int overflows after 12! (13! is larger than the max int value)
    }

    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}
func main() {
    // Test cases
    fmt.Println(FactorialIterative(6))   // Output: 24
    fmt.Println(FactorialIterative(0))   // Output: 1
    fmt.Println(FactorialIterative(-5))  // Output: 0
    fmt.Println(FactorialIterative(13)) // Output: 0 (overflow)
}