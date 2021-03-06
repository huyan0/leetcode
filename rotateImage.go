func rotate(matrix [][]int)  {
    // m'(i,j) = m(n-1-j, i)
    n := len(matrix)
    for i := 0; i < (n+1) / 2; i ++ {
        for j := i; j < n-1-i; j ++ {
            x, y, temp := i, j, matrix[i][j]
            // four numbers is a group. start at i, j
            // the last number should be temp
            for {
                if n-1-y==i && x == j {
                    matrix[x][y] = temp
                    break
                }
                matrix[x][y] = matrix[n-1-y][x]
                x, y = n-1-y, x
            }
            
        }
    }
}
