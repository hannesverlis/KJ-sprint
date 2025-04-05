package sprint


func Overlap(arr1, arr2 []int) []int {
	// Create a map to store the frequency of elements in arr1
	freq1 := make(map[int]int)
	for _, num := range arr1 {
		freq1[num]++
	}

	// Create a map to store the frequency of elements in arr2
	freq2 := make(map[int]int)
	for _, num := range arr2 {
		freq2[num]++
	}

	// Create a slice to store the common elements
	var common []int
	for num, count1 := range freq1 {
		if count2, exists := freq2[num]; exists {
			// Get the minimum count of the common element
			minCount := min(count1, count2)
			for i := 0; i < minCount; i++ {
				common = append(common, num)
			}
		}
	}
    if len(common) == 0 {
		return []int{}
	}
	// Return the common elements in sorted order
	return mergeSort(common)
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Merge Sort implementation to sort the array
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// Helper function to merge two sorted arrays
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}