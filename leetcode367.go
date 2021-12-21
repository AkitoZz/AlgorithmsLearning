package list

// 标准二分思想, 不断二分寻找中值是否为对应平方数, 根据大小移动左右边界
func isPerfectSquare(num int) bool {
    var left = 1
    var right = num
    var middle = 0

    for left <= right {
        middle = (left + right) / 2
        if middle * middle > num {
            right = middle - 1
        } else if middle * middle < num {
            left = middle + 1
        } else {
            return true
        }
    }

    return false
}
