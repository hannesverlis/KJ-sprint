package sprint

//array 8 / osa 4  //tehtud kl 22:46
func FilterBySum(arr [][]int, limit int) [][]int {
    filtered := make([][]int, 0) // Initialize the filtered array

    for _, subArr := range arr {
        sum := 0 

        // Calculate the sum of the current subarray
        for _, num := range subArr {
            sum += num
        }

        // If the sum meets the limit, add the subarray to the filtered slice
        if sum >= limit {
            filtered = append(filtered, subArr)
        }
    }
    return filtered
}


// FilterBySum([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, 9)
//>> //[[2 3 4] [3 4 5]]
