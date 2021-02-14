// Non-negative numbers
// (0,0) -> len(grid), len(grid[0]) (m-1, n - 1)
// can only move down or move right -> each node has at most two neighbours

import "fmt"
import "math"

func minPathSum(grid [][]int) int {
    
    // get m and n
    n := len(grid)
    
    if n < 1 {
        return -1
    }
    
    m := len(grid[0])
    dist := make([]([]int), n, m*n)
    
    for i, _ := range dist {
        dist[i] = make([]int, m)
        for j, _ := range dist[i] {
            dist[i][j] = math.MaxInt32
        }
    }
    
    visited := make([][]bool, n, m*n);
    
    for i, _ := range visited {
        visited[i] =  make([]bool, m)
    }
    
    dist[0][0] = grid[0][0]
    markAsVisited(0,0,visited)
    calcTempDist(0, 0, dist, grid)

    
    for !hasVisited(n-1, m-1, visited){
            curN, curM := findMinNode(dist, visited)
            markAsVisited(curN, curM, visited)
            calcTempDist(curN, curM, dist, grid)
    }
    fmt.Printf("%v", dist)
    return dist[n-1][m-1]
}


func markAsVisited(n,m int, visited [][]bool) {
    visited[n][m] = true;
} 

func hasVisited(n,m int, visited [][]bool) bool {
    return visited[n][m]
}

func findMinNode(dist [][]int, visited [][]bool) (n, m int){
    min := math.MaxInt32
    for i, arr := range dist {
        for j, num := range arr {
            if !hasVisited(i, j, visited) && num <= min {
                min = num
                n, m = i, j 
            }
        }
    }
    return
}

func calcTempDist(n, m int, dist, grid[][]int) {
    // down
    if n + 1 < len(dist)  && dist[n+1][m] > dist[n][m] + grid[n+1][m]{
        dist[n+1][m] = dist[n][m] + grid[n+1][m]
    }
    // right
    if m + 1 < len(dist[0]) &&  dist[n][m+1] > dist[n][m] + grid[n][m+1]{
        dist[n][m+1] = dist[n][m] + grid[n][m+1]
    }
}
