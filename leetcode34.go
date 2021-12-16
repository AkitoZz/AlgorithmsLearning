package list

// 升序，数据会重复，时间要求Ologn, 难度在于二分的边界值问题
// 在普通二分查找中，由于数据不重复，所以可以直接在找到目标值时返回数据下标。但由于该题有重复数据，所以要将直接return的逻辑改为继续移动下标寻找目标数据
// 在继续寻找数据的逻辑中，由于是升序数组，所以有以下几种情况：1.中间值非连续重复，即第一次命中的下标为起始下标（注意此时左中右三个边界值要依次判断，同时命中会影响顺序）
// 2.中间值为连续值，此时第一次命中的下标非起始下标，要去寻找连续数字的第一个下标作为起始下标
// 在上述两个大前提下即可找到满足条件的下标值，但是在处理结果时还有一些特殊情况：1.只有一个命中值（重复赋值即可） 2.超过两个命中值（取头尾）3.没有命中值（-1，-1）

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    var left = 0
    var right = len(nums) - 1
    var middle = 0
    var res []int

    for (left <= right) {
        middle = (left + right) / 2
        if target < nums[middle] {
            right = middle - 1
        } else if target > nums[middle] {
            left = middle + 1
        } else {
            if target == nums[left] {
                res = append(res, left)
                left = left + 1
                continue
            } else if target == nums[middle] {
                res = append(res, middle)
                //if len(res) >= 2 {
                //    break
                //}
                if len(res) == 1 && nums[middle - 1] == nums[middle] {
                    for i := 0; i < middle; i++ {
                        if target == nums[i] {
                            res[0] = i
                            break
                        }
                    }
                } else {
                    left = middle + 1
                }
                
                continue
            } else if target == nums[right] {
                res = append(res, right)
                right = right - 1    
                continue
            }
        }
    }

    if len(res) == 0 {
        return []int{-1, -1}
    } else if len(res) == 1 {
        res = append(res, res[0])
        return res
    } else if len(res) > 2 {
        return []int{res[0], res[len(res) - 1]}
    } else {
        return res
    }

}
