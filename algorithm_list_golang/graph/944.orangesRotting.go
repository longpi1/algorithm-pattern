package main

/*
994. 腐烂的橘子
在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

示例 1：
输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：
输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
示例 3：
输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
*/

func orangesRotting(grid [][]int) int {

}

/*
解题思路：
BPF
*/

var (
	/*
		dx = []int{1, 0, 0, -1} 表示在横坐标上，分别向右、不动、不动、向左移动；
		dy = []int{0, 1, -1, 0} 表示在纵坐标上，分别不动、向上、向下、不动移动。
	*/
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

// orangesRotting 使用广度优先搜索（BFS）算法解决腐烂橘子问题
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{}       // 用于存储腐烂橘子的坐标
	vis := make([][]bool, m) // 用于标记橘子是否被访问
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	good := 0 // 用于统计新鲜橘子的数量

	// 初始化，将腐烂橘子的坐标加入队列，并统计新鲜橘子的数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				good++
			}
		}
	}

	// 如果没有新鲜橘子或者没有腐烂橘子，返回相应的结果
	if good == 0 {
		return 0
	}
	if len(queue) == 0 {
		return -1
	}

	level := 0 // 用于记录腐烂的时间，即分钟数
	for len(queue) != 0 {
		p := [][]int{} // 用于存储新一轮腐烂的橘子坐标

		for i := 0; i < len(queue); i++ {
			node := queue[i]
			x := node[0]
			y := node[1]
			if grid[x][y] == 1 || vis[x][y] {
				continue
			}
			vis[x][y] = true

			// 对当前腐烂的橘子的四个方向进行遍历
			for idx := 0; idx < 4; idx++ {
				nx := x + dx[idx]
				ny := y + dy[idx]

				// 如果新坐标在网格内，且对应的橘子是新鲜的，则将其标记为腐烂，并加入新一轮腐烂的队列
				if nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					p = append(p, []int{nx, ny})
				}
			}
		}

		queue = p
		level++ // 进入下一分钟
	}

	// 如果仍然存在新鲜橘子，则返回-1，表示无法使所有橘子腐烂
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	// 返回腐烂所需的分钟数（减1是因为最后一轮腐烂结束时level会多加1）
	return level - 1
}

type orange struct {
	x int
	y int
}

var directions = []orange{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
	// 统计新鲜橘子数量 ，栈储存坏橘子
	fresh := 0
	q := []orange{}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				q = append(q, orange{i, j})
			}
		}
	}
	// 循环，直到都不新鲜或者不能腐烂
	res := 0
	for fresh > 0 && len(q) > 0 {
		res++
		tmp := q
		q = []orange{}
		for _, t := range tmp { // 已经腐烂橘子
			for _, d := range directions {
				i, j := t.x+d.x, t.y+d.y
				if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
					fresh--
					grid[i][j] = 2
					q = append(q, orange{i, j})
				}
			}
		}
	}
	if fresh > 0 {
		return -1
	}
	return res
}
