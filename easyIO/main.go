package main

import (
    "fmt"
    // "io"
)

func main() {
    // var a, b int
    // for {
    //     _, err := fmt.Scan(&a, &b)
    //     if err == io.EOF {
    //         break
    //     }
    //     fmt.Println(a + b)
    // }
	nums := []int{2, 8, 9}
    for i, v := range nums {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println()
	// 输出 2 8 9
}
