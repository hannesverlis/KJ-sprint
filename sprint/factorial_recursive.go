package sprint
func FactorialRecursive(n int) int {

	    // Error handling: Factorial is not defined for negative numbers or overflows
		if n <= 0 {
			if n > 20{
				return 0
			}
		}	
		if n == 0 {
			return 1
		}
		if n == -1 {
		return 0
		}

		
		return n * FactorialRecursive(n - 1)

		//result := n * FactorialRecursive(n-1)

		//return result
}	