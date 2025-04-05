package sprint

//import "fmt"

// PayAmount function takes an amount and a slice of denominations
// and returns the denominations used to pay the exact amount.
func Payout(amount int, denominations []int) []int {
	// Result slice to store the denominations used
	result := []int{}

	// Find the maximum denomination first, and process in that order
	for {
		maxDenom := 0
		maxIndex := -1
		for i, denom := range denominations {
			if denom > maxDenom {
				
				//fmt.Println(denom)
				//fmt.Println(maxDenom)
				maxDenom = denom
			//	fmt.Println(maxDenom)
				maxIndex = i
				//fmt.Println(maxIndex)

			}
		}
		// If no more denominations to process, break the loop
		if maxIndex == -1 {
			break
		}

		// Remove the processed denomination from the list
		denominations = append(denominations[:maxIndex], denominations[maxIndex+1:]...)

		// Use as many of the current denomination as possible
		for amount >= maxDenom {
			amount -= maxDenom
			result = append(result, maxDenom)
		}
	}

	// If we were able to pay the exact amount, return the result
	// Otherwise, return an empty slice
	if amount == 0 {
		return result
	}
	return []int{}
}