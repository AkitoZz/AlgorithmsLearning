package dfs

// 求子集, 回溯 & 枚举


// 枚举思路，每个数字都有是否存在于子集中两种状态，那么可以用二进制进行编码
// 假如给定数组[1,2,3] 那么000表示三个数字都不存在，001表示仅存在3,111表示三个数字都存在，将所有可能枚举即0 ~ 2^n - 1(n=len(num)),而这样的二进制编码可以视作一个掩码
// 和数组对应索引的位相与，为0表示不存在，为1表示存在
// 时间复杂度为O(n * 2^n) 2^n种排列，每一个掩码需要O(n)
func subsets(nums []int) [][]int {
    var mask = 0
    var length = len(nums)
    var res [][]int
    for ;mask < 1<<length; mask++ {
        tmp := []int{}
        for idx, val := range nums {
            if mask>>idx & 1 > 0 {
                tmp = append(tmp, val)
            }
        }
        res = append(res, tmp)
    }
    return res
}
