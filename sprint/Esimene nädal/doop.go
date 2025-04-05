package sprint

func Doop(a int, op string, b int) int {
	// Check for invalid number of arguments
	if len(op) != 1 { // Single character operators  // len - kontrollib operaatorite arvu.  != not equal
		return 0
	}

	// Handle division and modulo by zero
	if (op == "/" || op == "%") && b == 0 {
		return 0
	}

	// Perform the operation
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		return 0 // Invalid operator
	}
}
