package main

import (
	"fmt"
)

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]


提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9
*/

/*func plusOne(digits []int) []int {
	n := len(digits)
	var count int
	result := make([]int,n)

	for i := 0; i < n ; i++ {
		count = count + digits[i] *  int(math.Pow(10,float64(n-i -1)))
	}
	count = count + 1
	for i := 0; i <n ; i++ {
		size := count / int(math.Pow(10,float64(n-i -1)))
		count = count - size * int(math.Pow(10,float64(n-i -1)))
		if size == 10 {
			result[i] = 0
			result[i+1] = 1
			continue
		}
		result[i] = size
	}
	return result
}*/

/*
上述代码错误，错误原因如下：
1.未考虑清楚末尾为9的情况
*/

/*
正确思路：
当我们对数组 digits 加一时，我们只需要关注 digits 的末尾出现了多少个 999 即可。我们可以考虑如下的三种情况：

如果 digits 的末尾没有 999，例如 [1,2,3][1, 2, 3][1,2,3]，那么我们直接将末尾的数加一，得到 [1,2,4][1, 2, 4][1,2,4] 并返回；

如果 digitss 的末尾有若干个 999，例如 [1,2,3,9,9][1, 2, 3, 9, 9][1,2,3,9,9]，那么我们只需要找出从末尾开始的第一个不为 999 的元素，即 333，将该元素加一，得到 [1,2,4,9,9][1, 2, 4, 9, 9][1,2,4,9,9]。随后将末尾的 999 全部置零，得到 [1,2,4,0,0][1, 2, 4, 0, 0][1,2,4,0,0] 并返回。

如果 digits\textit 的所有元素都是 999，例如 [9,9,9,9,9][9, 9, 9, 9, 9][9,9,9,9,9]，那么答案为 [1,0,0,0,0,0][1, 0, 0, 0, 0, 0][1,0,0,0,0,0]。我们只需要构造一个长度比 digits\textit{digits}digits 多 111 的新数组，将首元素置为 111，其余元素置为 000 即可。

算法

们只需要对数组 digits\textit{digits}digits 进行一次逆序遍历，找出第一个不为 999 的元素，将其加一并将后续所有元素置零即可。如果 digits\textit{digits}digits 中所有的元素均为 999，那么对应着「思路」部分的第三种情况，我们需要返回一个新的数组。
*/

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}





func main(){
	digits := []int{9, 9}
	fmt.Printf("result: %v", plusOne(digits))
}