package sprint



func ArrMap(f func(int) bool, a []int) []bool {
    result := make([]bool, len(a)) // Create a slice to store results

    for i, num := range a {
        result[i] = f(num) // Apply the function to each element
    }

    return result
}


func IsPrime(n int) bool {
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

func IsNegative(n int) bool {
	return n < 0

}
