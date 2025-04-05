package sprint

func LCM(a, b int) int {
    // Handle negative input by converting to positive
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -b
    }

    // Calculate GCD (without Abs)
    gcd := GCD(a, b)

    // Formula: LCM(a, b) = (a * b) / GCD(a, b)
    return (a * b) / gcd
}

func GCD(a, b int) int {
    // Handle negative input by converting to positive
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -b
    }

    // Base case: if either number is zero, the GCD is the other number
    if a == 0 {
        return b
    }
    if b == 0 {
        return a
    }

    // Euclidean algorithm: repeatedly find the remainder until it's zero
    for b != 0 {
        a, b = b, a%b
    }

    return a
}