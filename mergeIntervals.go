// Intuition: sort by the start of the interval, then run through the sorted list
// one time to merge overlaping intervals. \
// This works because if intervals[i] does not overlapse with interval[j], i < j, and 
// the intervals are sorted, then intervals[i] will not overlapse with any interval[k],
// where k > j.
// The complexity of this is O(NlogN)

import "sort"

func merge(intervals [][]int) [][]int {
    
    sort.SliceStable(intervals, func(i, j int ) bool{
        return intervals[i][0] < intervals[j][0]
    })
    
    result := make([][]int, 0)
    if len(intervals) == 0 {
        return result
    }
    result = append(result, intervals[0])
    
    for _, nums := range intervals {
        // merge
        if nums[0] <= result[len(result) - 1][1] {
            result[len(result) - 1][1] = maxOf(result[len(result) - 1][1], nums[1])
        } else {
            result = append(result, nums)
        }
    }
    
    return result
}


func maxOf(nums ...int) int {
    max := nums[0]
    
    for _, num := range nums {
        if num > max {
            max = num
        }
    }
    return max
}
