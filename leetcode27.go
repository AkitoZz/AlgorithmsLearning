package list

// 思路同26，使用快慢指针，快指针用作遍历数组，并拿到目的数据，慢指针用作修改数组，每填入一个目的数据地址+1

func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }

    slow, fast := -1, 0
    for ;fast < len(nums); fast++ {
        if nums[fast] != val {
            slow++
            nums[slow] = nums[fast]
        }
    }

    return slow + 1
}
