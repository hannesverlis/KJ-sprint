package sprint

func TransposeMatrix(matrix [][]int) [][]int {
    // Check if the matrix is empty
    if len(matrix) == 0 {
        return matrix
    }

    // Determine dimensions of the transposed matrix
    rows, cols := len(matrix), len(matrix[0])

    // Create the transposed matrix
    transposed := make([][]int, cols)
    for i := range transposed {
        transposed[i] = make([]int, rows)
    }

    // Populate the transposed matrix
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            transposed[j][i] = matrix[i][j]
        }
    }

    return transposed
}