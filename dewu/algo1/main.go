// 题目描述：
// 给你一个只包含数字字符、'@'和'#'的字符串，表示一个表达式。

// 定义运算 a@b=a+b, a#b=a-(a&b)。其中 a&b 表示a和b二进制下按位与的结果。且 # 的优先级大于 @。

// 请你计算输入的表达式的值，保证表达式中的数字大小都不超过99999（不会有连续6个或以上的数字字符）且不含前导0。同时，可以证明，运算过程中不会有负数。

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// 定义运算符的优先级
var precedence = map[rune]int{
	'@': 1,
	'#': 2,
}

// applyOp 函数根据运算符计算结果
func applyOp(op rune, b, a int) int {
	switch op {
	case '@':
		return a + b
	case '#':
		// a#b = a - (a&b)
		return a - (a & b)
	}
	return 0
}

// Evaluate 函数是表达式求值的核心逻辑
func Evaluate(expression string) int {
	// 操作数栈
	var operands []int
	// 运算符栈
	var operators []rune

	for i := 0; i < len(expression); i++ {
		// 当前字符
		char := rune(expression[i])

		// 如果是数字，则解析完整数字
		if unicode.IsDigit(char) {
			numStr := ""
			for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
				numStr += string(expression[i])
				i++
			}
			i-- // 回退一位，因为外层 for 循环会 i++
			num, _ := strconv.Atoi(numStr)
			operands = append(operands, num)
		} else { // 如果是运算符
			// 当运算符栈不为空，且栈顶运算符优先级 >= 当前运算符优先级时，进行计算
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[char] {
				// 弹出运算符
				op := operators[len(operators)-1]
				operators = operators[:len(operators)-1]

				// 弹出两个操作数
				val2 := operands[len(operands)-1]
				operands = operands[:len(operands)-1]
				val1 := operands[len(operands)-1]
				operands = operands[:len(operands)-1]

				// 计算结果并压回操作数栈
				operands = append(operands, applyOp(op, val2, val1))
			}
			// 将当前运算符压入运算符栈
			operators = append(operators, char)
		}
	}

	// 处理栈中剩余的所有运算符
	for len(operators) > 0 {
		op := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		val2 := operands[len(operands)-1]
		operands = operands[:len(operands)-1]
		val1 := operands[len(operands)-1]
		operands = operands[:len(operands)-1]

		operands = append(operands, applyOp(op, val2, val1))
	}

	// 最终结果就是操作数栈中唯一剩下的元素
	return operands[0]
}

// main 函数用于演示和测试
func main() {
	// --- 示例 1 ---
	input1 := "3#5@2#4"
	fmt.Printf("输入表达式: %s\n", input1)
	result1 := Evaluate(input1)
	fmt.Printf("计算结果: %d\n", result1)
	// 逐步分析:
	// 1. 3#5 -> 3 - (3&5) = 3 - 1 = 2
	//    表达式变为: 2@2#4
	// 2. 2#4 -> 2 - (2&4) = 2 - 0 = 2
	//    表达式变为: 2@2
	// 3. 2@2 -> 2 + 2 = 4
	// 最终结果: 4

	fmt.Println("\n--------------------\n")

	// --- 示例 2 ---
	input2 := "10@20#30@40"
	fmt.Printf("输入表达式: %s\n", input2)
	result2 := Evaluate(input2)
	fmt.Printf("计算结果: %d\n", result2)
	// 逐步分析:
	// 1. 20#30 -> 20 - (20&30) = 20 - 20 = 0
	//    表达式变为: 10@0@40
	// 2. 10@0 -> 10 + 0 = 10
	//    表达式变为: 10@40
	// 3. 10@40 -> 10 + 40 = 50
	// 最终结果: 50

	fmt.Println("\n--------------------\n")

	// --- 示例 3 ---
	input3 := "123#456@789"
	fmt.Printf("输入表达式: %s\n", input3)
	result3 := Evaluate(input3)
	fmt.Printf("计算结果: %d\n", result3)
	// 逐步分析:
	// 1. 123#456 -> 123 - (123&456) = 123 - 72 = 51
	//    表达式变为: 51@789
	// 2. 51@789 -> 51 + 789 = 840
	// 最终结果: 840
}