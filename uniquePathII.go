func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
        return 0
    }
    
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([][]int, m)
    
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = 1
    for k := 1; k <= m + n - 2; k++ {
        for i := 0; i < m; i++ {
            j := k - i;
            if j >= 0 && j < n {
                if obstacleGrid[i][j] == 1 {
                    continue
                }
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
