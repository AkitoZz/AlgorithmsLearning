package list

// 二分查找, 前提：有序数组, 无重复元素
// 注意边界问题, 保持一个不变的循环条件, 假设左右下表为闭区间, 即left, right作为边界值也是有意义的
func search(nums []int, target int) int {
    var left = 0
    var right = len(nums) - 1
    var middle = 0
  
    for (left <= right) {
        middle = (left + right) / 2
        if nums[middle] > target {
            right = middle - 1
        } else if nums[middle] < target {
            left = middle + 1
        } else {
            return middle
        }
    }
    return -1
}
