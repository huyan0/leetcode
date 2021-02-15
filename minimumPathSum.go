// Non-negative numbers
// (0,0) -> len(grid), len(grid[0]) (m-1, n - 1)
// can only move down or move right -> each node has at most two neighbours
// use dp memo

func minPathSum(grid [][]int) int {
    
    // get m and n
    n := len(grid)
    
    if n < 1 {
        return -1
    }
    
    m := len(grid[0])
    
    for i, _ := range dist {
        dist[i] = make([]int, m)
    }
    
    // initialize starting position
    dist[0][0] = grid[0][0]
    
    // initialize column zero
    for i := 1; i < len(dist); i++ {
        dist[i][0] = dist[i-1][0] + grid[i][0]
    }
    // initialize row zero 
    for i := 1; i < len(dist[0]); i++ {
        dist[0][i] = dist[0][i-1] + grid[0][i]
    }
    
    for row := 1; row < n; row++ {
        for col := 1; col < m; col++ {
            dist[row][col] = minOf(dist[row-1][col], dist[row][col-1]) + grid[row][col]
        } 
    }
    
    
    return dist[n-1][m-1]
}

func minOf(nums  ...int) int{
    min := nums[0]
    for _, num := range nums {
        if num < min {
            min = num
        }
    }
    return min
}
