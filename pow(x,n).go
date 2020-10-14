// naive intuition: 
// Multiple x by itself n times using a loop. Since x^n cannot overflow, this problem seems 
// straight-forward? 
// For negative n, divide 1 by the product.
// The time complexity is O(N), and space complexity is O(1)
//
// So the naive intuition gave me Time Limit Exceeded. Gotta think of some optimization...
// I bet its something like bit shifting with respect to power of two...
//
// Intuition 2: since x^(2^m) is shifting x to the right by m bits, which reduces the cost of operation
// to O(1). If we can tranform n to 2^m, we can take advantage of this trick. We can get such m with
// m = log2(n), which means we can rewrite the original equation to x^n = x^(2^m) 
// However, if n is not a exponent of 2, there will be a modular. To take that into account, let 
// l = n - 2^floor(log2(n)), and the original equation would be:
// x^(2^m) * x^l. The cost of the function would be O(l). 
// 
// Intuition 2 does not seem to work either, because it does not make sense to perform bitwise operation 
// on float type, which is what x is....
// 
// I looked at other people's solution. What they do is use divide and concquer: x^n = x ^ n / 2 * x ^n / 2
// This reduce the cost to O(logN).
// Another great learning experience 芜湖 :）


func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    
    product, negative := 1.0, false
    
    if n < 0 {
        n *= -1
        negative = true
    }
    
    product *= calculatePow(x, n)
    
    if negative {
        return 1 / product 
    }
    return product
}

func calculatePow(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    if n == 1 {
        return x
    }
    half := calculatePow(x, n / 2)
    ans := half * half
    if n % 2 != 0 {
        return ans * x
    }
    return ans
}
