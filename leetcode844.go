package list

// 操作字符串中的数据，想到双指针
// 快指针用于寻找退格符"#", 遇到退格符时慢指针地址往回移动一步(为0时不移动), 不为退格符时则将快指针地址的值赋给慢指针地址, 从而获取目的字符串
// 注意go中字符串为不可变变量, 所以操作字符串前将其先赋值给一个byte切片作为操作源
func backspaceCompare(s string, t string) bool {
    var slow, fast = 0, 0

    sb := []byte(s)
    tb := []byte(t)
    for ;fast < len(sb); fast++ {
        if sb[fast] != '#' {
            sb[slow] = sb[fast]
            slow++
        } else {
            if slow != 0 {
                slow--
            }
        }
    }
    s1 := string(sb[0 : slow])
    //fmt.Println(s1)

    slow, fast = 0, 0
    for ;fast < len(tb); fast++ {
        if tb[fast] != '#' {
            tb[slow] = tb[fast]
            slow++
        } else {
            if slow != 0 {
                slow--
            }
        }
    }
    t1 := string(tb[0 : slow])
    //fmt.Println(t1)

    return s1 == t1
}
