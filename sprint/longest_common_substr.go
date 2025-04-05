package sprint

func LongestCommonSubstr(str1, str2 string) string {
    m, n := len(str1), len(str2)
    dp := make([][]int, m+1) // Dynamic programming table
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    maxLength, endIndex := 0, 0 // Track the longest substring details

    // Build the DP table
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                if dp[i][j] > maxLength {
                    maxLength = dp[i][j]
                    endIndex = i // Track the end index for retrieval
                }
            }
        }
    }
   
    // Extract the substring
    if maxLength > 0 {
        return str1[endIndex-maxLength : endIndex]
    }
    return "" 
}