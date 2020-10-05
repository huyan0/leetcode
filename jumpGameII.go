func jump(nums []int) int {
    res, cur, i := 0, 0, 0
    for cur < len(nums) - 1 {
        res++
        prev := cur
        for i <= prev {
            cur = maxOf(cur, nums[i] + i)
            i++
        }
    }    
    return res
}

func maxOf(num1, num2 int) int {
    if num1 > num2 {
        return num1
    }
    return num2
}


