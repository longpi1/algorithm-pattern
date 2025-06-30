package main

import "fmt"

/*
72. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符

示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/

// 第一遍思路错误
func minDistanceError(word1 string, word2 string) int {

	n1 := len(word1)
	n2 := len(word2)

	// 创建一个DP（动态规划）表，大小为 (n1+1) x (n2+1)。
	// dp[i][j] 的含义是 word1 的前 i 个字符和 word2 的前 j 个字符的“最长公共子序列”的长度。
	// 这个定义是LCS算法的，而不是编辑距离算法的。
	result := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		result[i] = make([]int, n2+1)
	}

	// ---------------- 问题点：缺少DP表的初始化 ----------------
	// 虽然Go语言默认将int初始化为0，恰好是LCS算法需要的边界条件，
	// 但在标准的DP写法中，应该显式地初始化第一行和第一列为0，以表示空字符串与任何字符串的LCS长度都是0。
	// 这使得代码意图更清晰。对于“编辑距离”算法，这里的初始化是完全错误的。
	// 编辑距离的初始化应为：result[i][0] = i, result[0][j] = j。

	// 双重循环，填充DP表。
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			// 如果当前两个字符相等
			if word1[i-1] == word2[j-1] {
				// 那么LCS的长度就是它们各自前一个位置的LCS长度加1。
				// 这是LCS算法的正确递推公式。
				result[i][j] = result[i-1][j-1] + 1

			} else {
				// 如果当前两个字符不相等
				// 那么LCS的长度等于 "word1前i个字符与word2前j-1个字符的LCS" 和
				// "word1前i-1个字符与word2前j个字符的LCS" 中的较大值。
				// 这也是LCS算法的正确递推公式。
				result[i][j] = max(result[i-1][j], result[i][j-1])
			}
		}
	}
	// 循环结束后，result[n1][n2] 存储的是 word1 和 word2 的最长公共子序列的长度 (LCS_length)。

	// ---------------- 问题点：错误的最终结果计算 ----------------
	// 这里的计算逻辑是错误的。它计算的是 `max(n1, n2) - LCS_length`。
	// 这个结果没有明确的、标准的算法意义。
	//
	// - 如果想求“编辑距离”，需要使用完全不同的DP递推公式（涉及插入、删除、替换三种操作）。
	// - 如果想利用LCS的结果求一个相关问题，比如“最少删除字符数使两字符串相等”，
	//   其公式应该是 `n1 + n2 - 2 * result[n1][n2]`。
	//
	// 因此，这个返回值对于函数名 "minDistance" 来说是完全不正确的。
	if n1 > n2 {
		return n1 - result[n1][n2]
	}
	return n2 - result[n1][n2]

}

/*

1. 核心思想：动态规划 (Dynamic Programming)
我们要解决的问题是：“将 word1 转换为 word2 所需的最少操作次数是多少？”

这个问题可以被分解为更小的、相互关联的子问题。例如，要计算 "horse" -> "ros" 的距离，我们可以先思考如何计算 "hors" -> "ro" 的距离，
或者 "horse" -> "ro" 的距离等。这种特性非常适合使用动态规划来解决。

我们会创建一个二维数组（DP表），通常命名为 dp。

dp[i][j] 的含义至关重要：它表示 将 word1 的前 i 个字符转换成 word2 的前 j 个字符所需要的最少操作次数。

我们的最终目标就是求出 dp[len(word1)][len(word2)]。

2. 初始化 (The Crucial First Step)
初始化是为动态规划设置好“边界条件”或“起始状态”。我们来思考一下 dp 表的第一行和第一列的含义。

初始化第一行 dp[0][j]

含义：将一个空字符串（word1 的前0个字符）转换成 word2 的前 j 个字符需要多少次操作？
答案：只能通过 j 次插入操作来完成。例如，将 "" 转换成 "ros" 需要插入 'r', 'o', 's'，共3次。
公式：dp[0][j] = j
初始化第一列 dp[i][0]

含义：将 word1 的前 i 个字符转换成一个空字符串需要多少次操作？
答案：只能通过 i 次删除操作来完成。例如，将 "horse" 转换成 "" 需要删除 'h', 'o', 'r', 's', 'e'，共5次。
公式：dp[i][0] = i
dp[0][0] 的特殊情况

含义：将空字符串转换成空字符串。
答案：需要0次操作。这符合上述两个公式 dp[0][0]=0。
初始化可视化：

假设 word1 = "horse"，word2 = "ros"。dp 表初始化后如下：

      ""  r   o   s   (word2)
""    0   1   2   3
h     1
o     2
r     3
s     4
e     5
(word1)
3. 状态转移方程（填充表格的规则）
现在我们需要填充表格的其余部分。对于 dp[i][j]，我们考虑 word1 的第 i 个字符 (word1[i-1]) 和 word2 的第 j 个字符 (word2[j-1])。

有两种情况：

情况一：word1[i-1] == word2[j-1] (两个单词的当前末尾字符相同)

例如，计算 "ros" -> "os" 的距离。因为末尾的 's' 相同，我们不需要对它进行任何操作。问题就简化为了计算 "ro" -> "o" 的距离。
规则：如果字符相同，当前位置的编辑距离就等于它们各自前一个位置的编辑距离。
公式：dp[i][j] = dp[i-1][j-1]
情况二：word1[i-1] != word2[j-1] (两个单词的当前末尾字符不同)

这时，我们有三种选择来让 word1 的前 i 个字符变成 word2 的前 j 个字符：

替换 (Replace)：将 word1[i-1] 替换成 word2[j-1]。这个操作本身需要1步，此前的子问题是把 word1 的前 i-1 个字符变成 word2 的前 j-1 个字符。总成本 = dp[i-1][j-1] + 1。
删除 (Delete)：将 word1[i-1] 删除。这个操作需要1步，然后我们需要把 word1 的前 i-1 个字符变成 word2 的前 j 个字符。总成本 = dp[i-1][j] + 1。
插入 (Insert)：在 word1 的前 i 个字符后面插入 word2[j-1]。这个操作需要1步，此前的子问题是把 word1 的前 i 个字符变成 word2 的前 j-1 个字符。总成本 = dp[i][j-1] + 1。
规则：我们希望操作次数最少，所以在这三个选择中取最小值。

公式：dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
*/
// minDistance 函数计算将 word1 转换为 word2 所需的最少操作数（插入、删除、替换）。
// 这就是经典的“编辑距离”或“莱文斯坦距离”问题。
func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)

	// 创建一个 DP 表，dp[i][j] 表示 word1 的前 i 个字符转换成 word2 的前 j 个字符的最小距离。
	// 大小为 (n1+1) x (n2+1) 是为了处理空字符串的情况。
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	// === 初始化 ===
	// 初始化第一列：将 word1 的前 i 个字符变为空字符串，需要 i 次删除。
	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	// 初始化第一行：将空字符串变为 word2 的前 j 个字符，需要 j 次插入。
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}

	// === 填充DP表 ===
	// 从 dp[1][1] 开始，根据状态转移方程填充。
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			// 如果当前考察的两个字符相同
			// 注意：字符串索引是 i-1 和 j-1，因为 dp 表的索引比字符串索引大 1
			if word1[i-1] == word2[j-1] {
				// 无需操作，成本等于它们各自前一个状态的成本。
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 如果字符不同，需要在“替换”、“删除”、“插入”三种操作中选择成本最小的一个。
				// 替换成本: dp[i-1][j-1] + 1
				// 删除成本: dp[i-1][j] + 1  (从 word1 删除一个字符)
				// 插入成本: dp[i][j-1] + 1  (在 word1 插入一个字符)
				dp[i][j] = min2(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	// 最终答案存储在表的右下角，代表将整个 word1 转换为整个 word2 的成本。
	return dp[n1][n2]
}

// 一个辅助函数，用于求三个整数的最小值。
func min2(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Printf("result", minDistance(word1, word2))
}
