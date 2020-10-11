// Complexity: N step to generate a single permutation, N! total permutations. So O(N * N!)

import (
    "sort"
)
func permuteUnique(nums []int) [][]int {
    // with duplicates, a common approach seems to be sort and make sure the same value does not appear on the same position multiple times
    sort.Ints(nums)
    // result two dimenstial slices
    res := make([][]int, 0)
    // visited slice to mark positions
    visited := make([]bool, len(nums))
    // call permuteDFS to recursively build answers
    permuteDFS(&res, visited, nums,  []int{})
    return res
}

func permuteDFS(res *[][]int, visited []bool, nums []int, cur []int) {
    // reached the end, add to res
    if len(cur) ==  len(nums) {
        temp :=  append([]int{}, cur...)
        *res = append(*res, temp)
        return
    }
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
            continue
        }
        if !visited[i] {
            visited[i]= true
            cur = append(cur, nums[i])
            permuteDFS(res, visited, nums, cur)
            cur = cur[:len(cur)-1]
            visited[i] = false;
        }
    }
}
