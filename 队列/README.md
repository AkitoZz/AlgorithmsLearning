###用栈实现队列（leetcode232题）

####实现方式
    栈（stack）的特性为FILO即先入后出，队列（queue）的特性为FIFO先入先出，用两个栈来实现队列的特性，
    即一个栈用作输入（push操作），一个栈用作输出（pop和peek操作），其中栈由静态数组实现。
    
- e.g1：输入1,2,3，将其依此压入input栈，需要输出操作时则将input栈的内容压入output栈（相当于翻转），
此时从output栈执行pop操作出来的便是数字1，实现了FIFO。
- e.g2：承接e.g1的结果，此时若需要输出则继续操作output栈，若要输入则操作input栈，假设此时输入4,5，
则将其压入input栈，此时的状态为：input（4,5），output（3,2），若这时要进行输出操作则继续操作output栈（
输出2,3），当output栈为空时，再将intput内容压入output，此时input为空，output为（5,4），继续执行输出操作（
输出4,5），实现了FIFO

    