package list 

// 二分查找，从1到x不断取中间数做平方与x比较大小，同时调整边界，需要注意终止条件和返回结果，终止条件为left<=right，如果无法整数平方那么肯定会移动到想上取整处的结果，所以返回left-1
func mySqrt(x int) int {
    if x == 1 {
        return 1
    }
    
    var left = 1
    var right = x
    var middle = 0

    for left <= right {
        middle = (left + right) / 2
        if middle * middle < x {
            left = middle + 1
        } else if middle * middle > x {
            right = middle - 1
        } else {
            return middle
        }
    }

    return left - 1
}
