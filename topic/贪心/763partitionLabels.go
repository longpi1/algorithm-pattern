package main

/*
763. 划分字母区间
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。

注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

返回一个表示每个字符串片段的长度的列表。
示例 1：
输入：s = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。
示例 2：
输入：s = "eccbbbbdec"
输出：[10]
*/

/*
此题未做出来：
思路
由于同一个字母只能出现在同一个片段，显然同一个字母的第一次出现的下标位置和最后一次出现的下标位置必须出现在同一个片段。因此需要遍历字符串，得到每个字母最后一次出现的下标位置。
在得到每个字母最后一次出现的下标位置之后，可以使用贪心的方法将字符串划分为尽可能多的片段，具体做法如下。

要有首尾两个指针，确定了结尾指针，就能确定下一个切割的开始指针。 遍历字符串，如果已扫描部分的所有字符，
都只出现在已扫描的范围内，即可做切割。
maintain「已扫描的字符能去到的最远位置」，扫到这个位置就切割，切出的字符不会在之后出现。 更新开始指针，准备下一次切割。
一些变量
maxPos 一个map（用数组速度可能会快一点），记录每个字母对应的最远位置。
start 做切割的开始位置。
scannedCharMaxPos 已扫描的字符能去到的最远位置。

作者：笨猪爆破组
链接：https://leetcode.cn/problems/partition-labels/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func partitionLabels(S string) []int {
	//一个map（用数组速度可能会快一点），记录每个字母对应的最远位置。
	maxPos := map[byte]int{}
	for i := 0; i < len(S); i++ {
		maxPos[S[i]] = i
	}
	// 用于保存结果
	res := []int{}
	// 做切割的开始位置。
	start := 0
	// scannedCharMaxPos 已扫描的字符能去到的最远位置。
	scannedCharMaxPos := 0
	for i := 0; i < len(S); i++ {
		curCharMaxPos := maxPos[S[i]]
		if curCharMaxPos > scannedCharMaxPos {
			scannedCharMaxPos = curCharMaxPos
		}
		if i == scannedCharMaxPos {
			res = append(res, i-start+1)
			start = i + 1
		}
	}
	return res
}
