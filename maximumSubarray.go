// intuition: 
// DP: at each number, the number is either part of the largest subarray started before the number. If it not part of the largest subarray before i, but part of some other subarray, the sum of the other subarray cannot be more than the largest array, thus we do not need to consider it
// or the begining of a new subarray, so the sum at  
// F(i) = max()

import "math"
func maxSubArray(nums []int) int {
    curSum, res := math.MinInt32,  math.MinInt32
    
    for _, num := range nums {
        curSum = maxOf(curSum + num, num)
        res = maxOf(res, curSum)
    }
    return res
}

func maxOf(nums ...int) int {
    max := nums[0]
    for _, num :=range nums {
        if max < num {
            max = num
        }
    }
    return max
}
