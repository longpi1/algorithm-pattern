package main

import "fmt"

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ['1','1','1','1','0'],
  ['1','1','0','1','0'],
  ['1','1','0','0','0'],
  ['0','0','0','0','0']
]
输出：1

示例 2：
输入：grid = [
  ['1','1','0','0','0'],
  ['1','1','0','0','0'],
  ['0','0','1','0','0'],
  ['0','0','0','1','1']
]
输出：3
提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && grid[i][j] == '1' && (grid[i-1][j] == '0' && grid[i+1][j] == '0') || (grid[i][j-1] == '0' && grid[i][j+1] == '0') {

			}
			if grid[i][j] == '1' && (grid[i-1][j] == '0' && grid[i+1][j] == '0') || (grid[i][j-1] == '0' && grid[i][j+1] == '0') {
				result++
			}
		}
	}
	return result
}

/*
思路一：深度优先遍历 DFS
目标是找到矩阵中 “岛屿的数量” ，上下左右相连的 1 都被认为是连续岛屿。
dfs方法： 设目前指针指向一个岛屿中的某一点 (i, j)，寻找包括此点的岛屿边界。
从 (i, j) 向此点的上下左右 (i+1,j),(i-1,j),(i,j+1),(i,j-1) 做深度搜索。
终止条件：
(i, j) 越过矩阵边界;
grid[i][j] == 0，代表此分支已越过岛屿边界。
搜索岛屿的同时，执行 grid[i][j] = '0'，即将岛屿所有节点删除，以免之后重复搜索相同岛屿。
主循环：
遍历整个矩阵，当遇到 grid[i][j] == '1' 时，从此点开始做深度优先搜索 dfs，岛屿数 count + 1 且在深度优先搜索中删除此岛屿。
最终返回岛屿数 count 即可。

*/

func numIslands(grid [][]byte) int {
	// 定义 DFS 函数，用于搜索相邻的陆地
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0' // 标记当前位置为已访问
		// 递归搜索相邻的陆地
		dfs(grid, i+1, j)
		dfs(grid, i, j+1)
		dfs(grid, i-1, j)
		dfs(grid, i, j-1)
	}

	count := 0
	// 遍历二维网格
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++         // 增加岛屿计数
				dfs(grid, i, j) // 执行 DFS 搜索

			}
		}
	}
	return count
}

/*
思路二：广度优先遍历 BFS
主循环和思路一类似，不同点是在于搜索某岛屿边界的方法不同。
bfs 方法：
借用一个队列 queue，判断队列首部节点 (i, j) 是否未越界且为 1：
若是则置零（删除岛屿节点），并将此节点上下左右节点 (i+1,j),(i-1,j),(i,j+1),(i,j-1) 加入队列；
若不是则跳过此节点；
循环 pop 队列首节点，直到整个队列为空，此时已经遍历完此岛屿。
*/
func numIslands(grid [][]byte) int {
	// 定义 BFS 函数，用于搜索相邻的陆地
	bfs := func(grid [][]byte, i, j int) {
		queue := [][]int{{i, j}}
		for len(queue) > 0 {
			// 出队列
			node := queue[0]
			queue = queue[1:]
			i, j := node[0], node[1]

			// 检查当前位置是否为陆地
			if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) && grid[i][j] == '1' {
				// 标记当前位置为已访问
				grid[i][j] = '0'
				// 将相邻的陆地位置加入队列
				queue = append(queue, []int{i + 1, j}, []int{i - 1, j}, []int{i, j - 1}, []int{i, j + 1})
			}
		}
	}

	count := 0
	// 遍历二维网格
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				continue // 如果当前位置是水域，跳过
			}
			count++         // 增加岛屿计数
			bfs(grid, i, j) // 执行 BFS 搜索

		}
	}
	return count
}

func main() {
	// 示例输入
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	result := numIslands(grid)
	fmt.Println("Number of islands:", result)
}
