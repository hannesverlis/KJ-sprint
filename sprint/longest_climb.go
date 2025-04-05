package sprint

func LongestClimb(arr []int) []int {
    if len(arr) == 0 {
        return []int{}
    }

    longestSubarray := []int{arr[0]} // Start with the first element
    currentSubarray := []int{arr[0]}
    
    for i := 1; i < len(arr); i++ {
        if arr[i] > arr[i-1] { 
            currentSubarray = append(currentSubarray, arr[i])
            
            if len(currentSubarray) > len(longestSubarray) {
                longestSubarray = currentSubarray
            }
        } else {
            currentSubarray = []int{arr[i]} // Reset if not increasing
        }
    }

    return longestSubarray
}