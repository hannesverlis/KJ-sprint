package sprint

func ToThePowerRecursive(base, exponent int) int {
    // Base case: Any number to the power 0 is 1
    if exponent == 0 {
        return 1
    }

    // Handle negative powers
    if exponent < 0 {
        return 0 
    }

    // Recursive case: x^n = x * x^(n-1)
    return base * ToThePowerRecursive(base, exponent-1)
}
