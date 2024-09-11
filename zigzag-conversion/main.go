package zigzag_conversion

import "strings"

/**
https://leetcode.cn/problems/zigzag-conversion/?envType=study-plan-v2&envId=top-interview-150

6. Z 字形变换

将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"

*/

// https://leetcode.cn/problems/zigzag-conversion/solutions/21610/zzi-xing-bian-huan-by-jyd/?envType=study-plan-v2&envId=top-interview-150
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	row := 0
	flag := -1
	for _, ch := range s {
		rows[row] += string(ch)
		// 第一行++ 最后一行--
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		row += flag
	}

	return strings.Join(rows, "")
}

// https://leetcode.cn/problems/zigzag-conversion/solutions/260155/ji-jian-jie-fa-by-ijzqardmbd/?envType=study-plan-v2&envId=top-interview-150
/**
以 V 字型为一个循环, 循环周期为 n = (2 * numRows - 2) （2倍行数 - 头尾2个）。
对于字符串索引值 i，计算 x = i % n 确定在循环周期中的位置。
则行号 y = min(x, n - x)。
*/
func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	// 循环周期 n = (2 * numRows - 2)
	n := numRows*2 - 2

	for i, char := range s {
		x := i % n
		// 行号 = min(x, n - x)
		rows[min(x, n-x)] += string(char)
	}
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
