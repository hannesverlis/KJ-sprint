package sprint

//array 8 / osa 5

func SortIntegerTable(table []int) []int {
	// Handle empty or single-element slices
	if len(table) < 2 {
		return table
	}

	// Bubble Sort Implementation
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(table)-1; i++ {
			if table[i] > table[i+1] {
				// Swap elements
				table[i], table[i+1] = table[i+1], table[i]
				swapped = true
			}
		}
	}

	return table
}

//SortIntegerTable([]int{5, 4, 3, 2, 1, 0})
//>> [0 1 2 3 4 5]
