// Intuition: in all the index that can be reached from the current index, find the index 
// i where i +  nums[i] gives the furthest result. Go to i, set it as cur, and repeat the 
// process. If i <= 0  and i + nums[i] < num[lenght -1] and i != length - 1, return false
func canJump(nums []int) bool {
    
    i, dest := 0, len(nums) - 1
    
    for i < dest && nums[i] > 0 {
        largestRange, nextIndex := 0, i
        
        for j := i+1; j <= i+nums[i] && j <= dest; j++ {
            if largestRange < j + nums[j] {
                largestRange, nextIndex = j + nums[j], j
            }   
        }
        
        i = nextIndex
    }
    return i >= dest
}
