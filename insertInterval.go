import "fmt"
// Grand Yang Solution: find the first overlapse, update newInterval to incorporate overlaped
// intervals, find the next non-overlapping interval, add to res 
func insert(intervals [][]int, newInterval []int) [][]int {
    res, cur := make([][]int, 0, len(intervals)), 0
    
    for cur < len(intervals) && intervals[cur][1] < newInterval[0] {
        res = append(res, intervals[cur])
        cur++
    }
    for cur < len(intervals) && intervals[cur][0] <= newInterval[1] {
        newInterval[0] = minOf(intervals[cur][0], newInterval[0])
        newInterval[1] = maxOf(intervals[cur][1], newInterval[1])
        cur++
    }
    res = append(res, newInterval)
    
    if cur < len(intervals) {
        res = append(res, intervals[cur:]...)    
    }
    
    return res
}

func minOf(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxOf(a, b int) int {
    if a < b {
        return b
    }
    return a
}

// Home-made solution: insert the interval to the right position based on newInterval[0]
// then run through elements to merge with other intervals
/*
func insert(intervals [][]int, newInterval []int) [][]int {
    i := 0

    for i < len(intervals) {
        if intervals[i][0] > newInterval[0] {
            // the most succint way to insert an element into position i of a slice
            intervals = append(intervals[:i+1], intervals[i:]...)
            intervals[i] = newInterval
            break
        }
        i++
    }
    if i == len(intervals) {
        intervals = append(intervals, newInterval)
    }

    // check with the one in front of it
    if i != 0 && intervals[i-1][1] >= newInterval[0]{
        intervals[i-1][1] = maxOf(intervals[i-1][1], newInterval[1])
        // remove new intervals from the list
        if i == len(intervals) - 1 {
            intervals = intervals[:i]
        } else {
            intervals = append(intervals[:i], intervals[i+1:]...)    
        }
        i--
    }
    
    for i < len(intervals) - 1 {
        // overlapse, merge
        if intervals[i][1] >= intervals[i+1][0] {
            intervals[i+1][0] = intervals[i][0] 
            intervals[i+1][1] = maxOf(intervals[i][1], intervals[i+1][1])
            intervals = append(intervals[:i], intervals[i+1:]...)
        } else {
            break
        }
    }
    return intervals
    
}
*/
