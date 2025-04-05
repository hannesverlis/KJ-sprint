package sprint




// AdvancedSortWordArr sorts a string slice using the provided comparison function.
func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
    sortedArr := make([]string, len(a)) // Create a copy of the original slice
    copy(sortedArr, a)                 // Copy the original data to avoid modifying it

    // Bubble Sort Implementation
    n := len(sortedArr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if f(sortedArr[j], sortedArr[j+1]) > 0 {
                // Swap elements if they are in the wrong order
                sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
            }
        }
    }

    return sortedArr
}
func StrCompare(a , b string) int {

	for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] < b[i] {
            return -1
        } else if a[i] > b[i] {
            return 1
        }
    }
	return 0
}