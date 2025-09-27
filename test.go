package main

import "fmt"

func main() {
    str := "1563693962"
    var stack []byte
    var maxStack []byte
    num := 1 // 当前递增序列长度，初始化为1因为第一个元素算长度1
    max := 1 // 最长递增序列长度，初始化为1

    // 将第一个字符放入栈
    stack = append(stack, str[0]-'0')
    
    for i := 1; i < len(str); i++ {
        v := str[i] - '0' // 转为数字
        if v > stack[len(stack)-1] { // 如果当前数字大于栈顶
            stack = append(stack, v)
            num++
            if num > max { // 更新最长序列
                max = num
                maxStack = make([]byte, len(stack))
                copy(maxStack, stack) // 保存当前最长序列
            }
        } else { // 当前数字不大于栈顶，重置栈
            stack = stack[:0] // 清空栈
            stack = append(stack, v) // 放入当前数字
            num = 1 // 重置当前长度为1
        }
    }

    // 输出结果
    fmt.Printf("最长递增子序列长度: %d\n", max)
    fmt.Printf("最长递增子序列: %v\n", maxStack)
}