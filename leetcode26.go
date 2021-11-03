package list

// leetcode 26 有序数组去重
// 值得注意的是传入的是数组指针，返回的是数组长度，即将单一元素放到对应长度即可
// 有序数组，去掉重复元素，思路：用快慢指针，慢指针用作指向目标地址，因为是要去重数组，所以每找一个不重复的元素的慢指针便记录该值并向后移动一步，
// 快指针用作找到重复元素的最后一个值，并将其写到慢指针地址，后续继续去找下一个重复元素的最后一个值，直到遍历完整个数组

// e.g [1,2,3,4,4,5,5]
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    slow, fast := -1, 0
    for ; fast <= len(nums) - 1; fast++ {
        if fast == 0 || nums[slow] != nums[fast] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}
