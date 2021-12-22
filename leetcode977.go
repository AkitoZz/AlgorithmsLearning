package list

// 运用双指针有两种算法
// 1. 头尾指针, 依次将最大的数据倒序放入目标数组中即可获得目标数组, 因为平方后负数部分是一个降序, 正数部分是一个升序, 这其实类似一个倒过来的归并排序
// 2. 直接使用归并排序, 找到正负数的分界线, 平方后获得两个数组, 负数部分是降序, 正数部分是升序, 负数数组倒序和正数部分做归并, 依次取出较小的那个数组成一个新数组
func sortedSquares(nums []int) []int {
    var left = 0
    var right = len(nums) - 1
    var res = make([]int, len(nums))
    var pos = len(nums) - 1

    for ;pos >= 0; pos-- {
        l := nums[left] * nums[left]
        r := nums[right] * nums[right] 

        if l > r {
            res[pos] = l
            left++
        } else {
            res[pos] = r
            right--
        }
    }

    return res
}
