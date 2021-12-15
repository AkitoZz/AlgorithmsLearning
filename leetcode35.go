package list


// 有序数组，无重复元素，查找数据，无则插入，要求O(log(n)) ，所以想到二分法
// 关键：在使用二分查找时，由于每次寻找的中间值坐标为向下取整，所以会导致如果目标数据在左侧时，会比预期向左偏移一个单位，可以通过middle和right的关系判断是否需要修正
func searchInsert(nums []int, target int) int {
    if target < nums[0] {
        return 0
    }
    if target > nums[len(nums) - 1] {
        return len(nums)
    }

    var left = 0
    var right = len(nums) - 1
    var middle = 0

    for (left <= right) {
        middle = (left + right) / 2
        if target > nums[middle] {
            left = middle + 1
        } else if target < nums[middle] {
            right = middle - 1
        } else {
            return middle
        }
        //fmt.Println(left, right ,middle)
    }

    if middle > right {
        return middle
    } else {
        return middle + 1
    }

}
