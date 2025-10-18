// 输入描述：每个测试文件均包含多组测试数据。第一行输入一个整数 $T (1 \le T \le 100)$ 代表数据组数，
// 每组测试数据描述如下：在第一行输入一个长度为 $1 \le length(n) \le 10^4$ 的十六进制数 $n$，且 $n$ 仅包含 A,B,C,D,E,F。
//  (注：根据样例推断，此处的描述应为“一个十六进制数n”，包含0-9和A-F)
//  输出描述：对于每组测试数据，新起一行，按顺序输出出现次数最多的八进制数字。若有多种数字并列出现次数最多，按升序依次输出它们，数字之间以单个空格分隔。
//  示例1输入：
// 2
// AA
// ABCD
// 输出：2
// 1 5

package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

// solve 函数处理单个测试用例
// 接收一个十六进制字符串，返回一个按要求格式化的八进制高频数字字符串
func solve(hexString string) string {
	// 1. & 2. 使用 math/big 进行超长十六进制到八进制的转换
	n := new(big.Int)
	// n.SetString(hexString, 16) 将16进制字符串解析为10进制存入 n
	// 题目保证了输入合法，所以我们忽略返回的 bool 值
	n.SetString(hexString, 16)

	// n.Text(8) 将 n 转换为 8 进制的字符串
	octalString := n.Text(8)

	// 3. 频率统计
	counts := make(map[rune]int)
	maxFreq := 0
	for _, r := range octalString {
		counts[r]++
		if counts[r] > maxFreq {
			maxFreq = counts[r]
		}
	}

	// 4. 收集频次最高的结果
	var resultRunes []rune
	for digit, freq := range counts {
		if freq == maxFreq {
			resultRunes = append(resultRunes, digit)
		}
	}

	// 5. 对结果进行升序排序
	sort.Slice(resultRunes, func(i, j int) bool {
		return resultRunes[i] < resultRunes[j]
	})

	// 6. 格式化输出
	var resultStrings []string
	for _, r := range resultRunes {
		resultStrings = append(resultStrings, string(r))
	}

	return strings.Join(resultStrings, " ")
}

// main 函数按照题目要求，模拟完整的输入和输出
func main() {
	// --- 模拟输入 ---
	// 题目中的示例输入
	T := 2
	inputs := []string{"AA", "ABCD"}

	// --- 完整的输入 ---
	// 按照您的要求，"自己输出完整的输入输出"
	// 我们先将模拟的输入打印出来
	fmt.Println("--- 完整的输入 ---")
	fmt.Println(T)
	for _, input := range inputs {
		fmt.Println(input)
	}

	// --- 完整的输出 ---
	// 接着，我们运行解题函数并打印输出
	fmt.Println("\n--- 完整的输出 ---")
	for _, input := range inputs {
		output := solve(input)
		fmt.Println(output)
	}

	/*
		// --- 在线评测(OJ)平台提交时，应使用以下标准输入(stdin)代码 ---
		var T int
		fmt.Scan(&T)
		for i := 0; i < T; i++ {
			var hexString string
			fmt.Scan(&hexString)
			fmt.Println(solve(hexString))
		}
	*/
}