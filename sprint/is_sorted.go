package sprint




func IsSorted(f func(a, b string) int, arr []string) bool {
    if len(arr) <= 1 {
        return true // An empty or single-element slice is always sorted
    }

    isAscending := f(arr[0], arr[1]) <= 0 // Detect initial sort order

    for i := 1; i < len(arr)-1; i++ {
        comparison := f(arr[i], arr[i+1])
        if (isAscending && comparison > 0) || (!isAscending && comparison < 0) {
            return false // Sort order violation
        }
    }
    return true
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