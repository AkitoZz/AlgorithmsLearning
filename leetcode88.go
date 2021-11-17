package list

// 有序数组合并排序
// 思路：双指针，从两个数组的头开始遍历，将小的数放进临时数组后指针向后偏移，需要指针走完两个数组，所以当一个数组遍历完时可以直接将另一个数组的后续部分挂到临时数组后端（因为两个目标数组是有序的）

func merge(nums1 []int, m int, nums2 []int, n int)  {
    var res []int
    var idxm, idxn = 0, 0
    // 终止条件为遍历完一个数组
    for idxm < m && idxn < n {
        if nums1[idxm] < nums2[idxn] {
            res = append(res, nums1[idxm])
            idxm++
        } else {
            res = append(res, nums2[idxn])
            idxn++
        }
    }

    if idxm == m {
        res = append(res, nums2[idxn:]...)
    } else {
        res = append(res, nums1[idxm:]...)
    }
    copy(nums1, res)
}
