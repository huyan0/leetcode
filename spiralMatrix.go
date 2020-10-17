func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    if m == 0 {
        return []int{}
    }
    
    n := len(matrix[0])
    if n == 0 {
        return []int{}
    }
    
    res := make([]int, 0, m*n)
    
    up, down, left, right := 0, m-1, 0, n-1
    
    for len(res) < m*n {
        // Right.
        for j := left; j <= right && len(res) < m*n; j++ {
            res = append(res, matrix[up][j])
        }
        
        // Down.
        for i := up+1; i <= down && len(res) < m*n; i++ {
            res = append(res, matrix[i][right])
        }
        
        // Left.
        for j := right-1; j >= left && len(res) < m*n; j-- {
            res = append(res, matrix[down][j])
        }
        
        // Right.
        for i := down-1; i >= up+1 && len(res) < m*n; i-- {
            res = append(res, matrix[i][left])
        }
        
        left++
        right--
        up++
        down--
    }
    
    return res
}
