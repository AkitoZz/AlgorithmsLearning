package list

// 双指针, 滑动窗口(时间复杂度On, 空间复杂度O1)
// 因为是要获取一个符合目标的连续数组的长度, 暴力解法为两层循环, 以一个数为基准, 获取它为开头的满足连续数组和大于目标值, 并判断这个长度是否为当前最小, 时间复杂度为On^2
// 滑动窗口思路, 双指针框出一个目标数组, 而这个框法是移动两个指针作为边界, 双指针均从起点开始移动, 不断移动右指针, 直到窗口大小大于目标值, 此时开始移动左指针, 尝试寻找更小的窗口, 没有则继续移动右指针
// 窗口滑动的关键点有两个
// 1. 只移动右指针, 这样窗口大小不用重复计算, 
// 2. 每次移动左指针时需要将窗口大小减去对应数值
// 这样看似是两层for循环, 其实关键在于外层循环为On, 内层循环每次操作复杂度是O1
func minSubArrayLen(target int, nums []int) int {
    var left, right = 0, 0
    var length = len(nums)
    var res = length + 1
    var sum = 0
    for ;right < length; right++  {
        sum += nums[right]
        
        for sum >= target {
            tmp := right - left + 1
            if tmp < res {
                res = tmp
            }
            sum -= nums[left]
            left++
        }    
    }

    if res == length + 1 {
        return 0
    }
    return res
}
