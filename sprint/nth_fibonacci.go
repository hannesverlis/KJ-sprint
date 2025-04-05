package sprint

func NthFibonacci(index int) int {
    // Base cases
    if index < 0 {
        return -1 // Handle negative indices
	}
    if index == 0 {
        return 0
}   
		if index == 1 {
        return 1
}
    // Recursive step
    return NthFibonacci(index-1) + NthFibonacci(index-2)
}