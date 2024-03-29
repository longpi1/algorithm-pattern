package main

import "fmt"

/*
滚动数组（Rolling Array），也称为滑动数组（Sliding Array）或循环数组（Circular Array），是一种用于优化内存使用的数据结构技巧。它的基本思想是在一个有限大小的数组或缓冲区中存储数据，当需要添加新数据时，新数据会覆盖掉旧数据，以保持数组的大小不变。
这种技巧通常用于处理动态规划问题中的状态转移，降低了内存消耗，尤其适用于只需要存储一部分之前的状态信息的问题。
滚动数组的优势在于避免了创建和维护大型的动态规划表格或二维数组，从而减少了内存占用。它通常用于那些只需要访问前一行或前一列数据的问题，而不需要保留整个表格的情况。通过滚动数组，只需要存储两个一维数组，一个用于当前层状态，
一个用于上一层状态，然后在迭代过程中交替更新这两个数组，以实现状态转移。
滚动数组在一些经典的动态规划问题中非常有用，例如0/1背包问题、斐波那契数列、最长公共子序列等。通过减少内存占用，它可以帮助提高程序的效率和性能，特别是在处理大规模问题或内存受限的情况下。

终于搞懂为啥要倒叙遍历了。
首先要明白二维数组的递推过程，然后才能看懂二维变一维的过程。
假设目前有背包容量为10，可以装的最大价值， 记为g(10)。
即将进来的物品重量为6。价值为9。
那么此时可以选择装该物品或者不装该物品。
如果不装该物品，显然背包容量无变化，这里对应二维数组，其实就是取该格子上方的格子复制下来，就是所说的滚动下来，直接g【10】 = g【10】，这两个g【10】要搞清楚，右边的g【10】是上一轮记录的，也就是对应二维数组里上一层的值，而左边是新的g【10】，
也就是对应二维数组里下一层的值。
如果装该物品，则背包容量= g(10-6) = g(4) + 9 ，也就是 g(10) = g(4) + 6 ,这里的6显然就是新进来的物品的价值，g(10)就是新记录的，对应二维数组里下一层的值，而这里的g(4)是对应二维数组里上一层的值，通俗的来讲：你要找到上一层也就是上一状态下
背包容量为4时的能装的最大价值，用它来更新下一层的这一状态，也就是加入了价值为9的物品的新状态。
这时候如果是正序遍历会怎么样？ g(10) = g(4) + 6 ，这个式子里的g(4)就不再是上一层的了，因为你是正序啊，g(4) 比g(10)提前更新，那么此时程序已经没法读取到上一层的g(4)了，新更新的下一层的g(4)覆盖掉了，这里也就是为啥有题解说一件物品被拿了两次的原因。
*/
func test_1_wei_bag_problem(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([]int,bagWeight+1)
	// 递推顺序
	for i := 0 ;i < len(weight) ; i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j:= bagWeight; j >= weight[i] ; j-- {
			// 递推公式
			dp[j] = max1(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Println(dp)
	return dp[bagWeight]
}

func max1(a,b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	weight := []int{1,3,4}
	value := []int{15,20,30}
	test_1_wei_bag_problem(weight,value,4)
}