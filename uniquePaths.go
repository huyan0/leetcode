
func uniquePaths(m int, n int) int {
    if n == 0 {
        return 0
    }
    
    dp := make([][]int, 0, m)
    
    for i := 0; i < m; i++ {
        dp = append(dp, make([]int, n))
    }
    
    dp[0][0] = 1
    for k := 1; k <= m + n - 2; k++ {
        for i := 0; i < m; i++ {
            j := k - i
            if j >= 0 && j < n {
                if i - 1 >= 0 {
                    dp[i][j] += dp[i-1][j]
                }
                
                if j - 1 >= 0 {
                    dp[i][j] += dp[i][j-1]
                }
            }
        }
    }
    return dp[m-1][n-1]
}
