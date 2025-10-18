// 小A喜欢看体育比赛。他对这次参赛的 $n$ 个选手都做了详细的调查，并对他们的实力进行了建模。
// 小A认为，第 $i$ 个选手拥有进攻能力 $a_i$ 与防守能力 $b_i$。当第 $i$ 个选手与第 $j$ 个选手对决时，
// 若 $a_i \times b_j > a_j \times b_i$，那么第 $i$ 个选手会获胜；若 $a_i \times b_j \le a_j \times b_i$，那么平凡，不分任何胜负。
// 小A已经提前建模好了每个选手的进攻能力和防守能力，他想用这个模型预测每个选手可以战胜多少人。

// 输入描述：第一行1个整数 $n$，表示选手的数量。第二行 $n$ 个整数 $a_1, a_2, ..., a_n$，表示选手的进攻能力。
// 第三行 $n$ 个整数 $b_1, b_2, ..., b_n$，表示选手的防守能力。$1 \le n \le 50000$, $0 \le a_i, b_i \le 100000$

// 输出描述：输出一行 $n$ 个数，分别表示小A的建模认为第 $i$ 个选手能战胜多少人。样例输入：3
// 3 2 1
// 2 1 1
// 样例输出：1 2 0

package main

import (
	"bufio"
	"fmt"
	// "os"
	"strconv"
	"strings"
)

// 为了在本地环境中方便地演示完整的输入和输出，
// 我们将使用 strings.NewReader 来模拟标准输入 (os.Stdin)。
// 在实际提交时，应该从 os.Stdin 读取。

// --- 模拟输入 ---
var mockInput = `
3
3 2 1
2 1 1
`

// ---------------------------------
// 在真实环境中提交时，请使用 os.Stdin
// var scanner = bufio.NewScanner(os.Stdin)
// ---------------------------------

// readInt 是一个辅助函数，用于从 scanner 中读取一个整数
func readInt(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		panic("Failed to scan integer")
	}
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Failed to convert text to integer")
	}
	return num
}

func main() {
	// --- 完整的输入 ---
	fmt.Println("--- 完整的输入 ---")
	fmt.Println(strings.TrimSpace(mockInput))
	fmt.Println("--------------------")

	// ---------------------------------
	// 真实提交时，请替换为 os.Stdin
	scanner := bufio.NewScanner(strings.NewReader(mockInput))
	// ---------------------------------
	
	// 使用 ScanWords 来按空格分割输入
	scanner.Split(bufio.ScanWords)

	// --- 1. 读取输入 ---
	n := readInt(scanner)

	// 使用 int64 来存储a和b，防止后续乘法溢出
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = int64(readInt(scanner))
	}

	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = int64(readInt(scanner))
	}

	// --- 2. 核心算法 O(n^2) ---
	// 遍历每个选手 i
	
	// 创建结果数组
	results := make([]int, n)
	
	for i := range n {
		count := 0
		// 遍历所有其他选手 j
		for j := 0; j < n; j++ {
			// 选手不和自己比较
			if i == j {
				continue
			}

			// 规则：ai * bj > aj * bi
			// 使用 int64 进行比较，防止溢出
			if a[i]*b[j] > a[j]*b[i] {
				count++
			}
		}
		results[i] = count
	}

	// --- 3. 完整的输出 ---
	fmt.Println("\n--- 完整的输出 ---")

	// 为了高效地拼接字符串，使用 strings.Builder
	var outputBuilder strings.Builder
	for i := 0; i < n; i++ {
		outputBuilder.WriteString(strconv.Itoa(results[i]))
		if i < n-1 {
			outputBuilder.WriteString(" ")
		}
	}
	
	// 打印最终结果
	fmt.Println(outputBuilder.String())
	
	// --- 4. 样例推演 ---
	fmt.Println("\n--- 样例推演 (a=[3 2 1], b=[2 1 1]) ---")
	fmt.Println("P1 (a=3, b=2):")
	fmt.Println("  vs P2 (a=2, b=1): 3*1 > 2*2 -> 3 > 4 (False)")
	fmt.Println("  vs P3 (a=1, b=1): 3*1 > 1*2 -> 3 > 2 (True)")
	fmt.Println("  -> P1 胜场: 1")
	fmt.Println("P2 (a=2, b=1):")
	fmt.Println("  vs P1 (a=3, b=2): 2*2 > 3*1 -> 4 > 3 (True)")
	fmt.Println("  vs P3 (a=1, b=1): 2*1 > 1*1 -> 2 > 1 (True)")
	fmt.Println("  -> P2 胜场: 2")
	fmt.Println("P3 (a=1, b=1):")
	fmt.Println("  vs P1 (a=3, b=2): 1*2 > 3*1 -> 2 > 3 (False)")
	fmt.Println("  vs P2 (a=2, b=1): 1*1 > 2*1 -> 1 > 2 (False)")
	fmt.Println("  -> P3 胜场: 0")
	fmt.Println("最终结果: 1 2 0")
}
