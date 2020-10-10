func permute(nums []int) [][]int {
    length := len(nums)
    
    res := make([][]int, 0)
    if length <= 1 {
        return [][]int{nums}
    }
    // We must return the slice afterwards because, although Append can modify the elements of slice, 
    // the slice itself (the run-time data structure holding the pointer, length, and capacity) is passed by value.
    for i, num := range nums {
        var intermediate [][]int
        var temp []int
        intermediate = permute(append(append(temp, nums[:i]...), nums[i+1:]...)) 
 
        for _, arr := range intermediate {
            res = append(res, append(arr, num))
        }
    }
    return res
}
