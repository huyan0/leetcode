func isMatch(s string, p string) bool {
    // recursive solution 
    // a-z and ? case easy to handle
    // * has different cases:
    // 1. * match to empty - move p forward by one, s stays the same
    // 2. * matches the first character in s - move s forward by one, p stays the same
    
    // DP solution: 
    // dp[len(s) + 1][len(p) + 1]
    // dp[0][0] = true
    // if p[j] == *
    //  dp[0][j] = dp[0][j-1] 
    // for i = 1 in s 
    //    for j n p 
    //    if
    
    // To find dp[i][j], we need : 
    // dp[i][j] = dp[i-1][j-1] if p[j] == ? || a-z, or 
    //            dp[i-1][j] => * matches s[i] || dp[i][j-1] => * is empty, if p[j] = *
    
    // initialize a len(x) + 1 by len(y) + 1 2D slice
    dp := make([][]bool, len(s) + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, len(p) + 1)
    }
    
    // set empty cases 
    dp[0][0] = true
    
    for i := 0; i < len(p); i++ {
        if p[i] == '*' {
            dp[0][i+1] = dp[0][i]
        } 
    }
    
    for i := 1; i < len(s) + 1; i++ {
        for j := 1; j < len(p) + 1; j++ {
            if p[j-1] == '*' {
                dp[i][j] = dp[i-1][j] || dp[i][j-1]
            } else if s[i-1] == p[j-1] || p[j-1] == '?' {
                dp[i][j] = dp[i-1][j-1] 
            }
        }
    }
    return dp[len(s)][len(p)]
    
}
