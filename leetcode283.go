package list

// 思路同26、27，双指针移动，慢指针用于填目的数据，填写一次向后移动一步，快指针用作寻找非零数据，找到后写入慢指针地址，并且将快指针地址填为0，因为后续慢指针会覆盖
// 思路为覆盖，交换的思路尚未考虑

func moveZeroes(nums []int)  {
    if len(nums) == 0 {
        return
    }

    slow, fast := -1, 0
    for ;fast < len(nums); fast++ {
        if nums[fast] != 0 {
            slow++
            nums[slow] = nums[fast]
            
            if fast > slow {
                nums[fast] = 0
            }
        }
    }
}
