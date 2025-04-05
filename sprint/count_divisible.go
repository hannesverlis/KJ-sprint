package sprint


func CountDivisible(from, to, step, divisor int) int {
    // Exception Handling
    if step <= 0 {
        //fmt.Println("Error: Step must be a positive integer.")
        return 0
    }

    if divisor == 0 {
        //fmt.Println("Error: Divisor cannot be zero.")
        return 0
    }

    count := 0

    // Loop through the range
    for i := from; i < to; i += step {
        if i%divisor == 0 {
            count++
        }
    }

    return count
}
