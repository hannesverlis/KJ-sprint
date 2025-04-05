package sprint

func PascalsTriangle(n int) [][]int {
    // Create the 2D slice to store the triangle
    triangle := make([][]int, n)

    // Generate each row of the triangle
    for i := 0; i < n; i++ {
        // Each row has i + 1 elements
        triangle[i] = make([]int, i+1)
        triangle[i][0] = 1   // First element is always 1
        triangle[i][i] = 1   // Last element is always 1

        // Calculate the middle elements
        for j := 1; j < i; j++ {
            triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
        }
    }
    return triangle
}