package main

import "fmt"

func main() {
    // 使用 make 创建切片
    s1 := make([]int, 3, 5) // 长度 3，容量 5
    fmt.Printf("s1: len=%d, cap=%d, %v\n", len(s1), cap(s1), s1)

    // 字面量创建切片
    s2 := []int{1, 2, 3} // 长度和容量相等
    fmt.Printf("s2: len=%d, cap=%d, %v\n", len(s2), cap(s2), s2)

    // 从数组或切片截取
    arr := [5]int{1, 2, 3, 4, 5}
    s3 := arr[1:3] // 从索引 1 到 3，容量从 1 到数组末尾
    fmt.Printf("s3: len=%d, cap=%d, %v\n", len(s3), cap(s3), s3)
}