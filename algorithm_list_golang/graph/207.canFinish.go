package main

/*
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
*/


func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses) // 用于保存课程之间的依赖关系，edges[i] 存储了课程 i 的后续课程列表
		indeg = make([]int, numCourses) // 记录每个课程的入度，即有多少课程依赖于该课程
		result []int // 用于存储拓扑排序的结果
	)

	// 构建图，初始化课程依赖关系和入度
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0]) // 课程 info[1] 依赖于课程 info[0]
		indeg[info[0]]++ // 课程 info[0] 的入度增加
	}

	q := []int{} // 用于存储入度为 0 的课程的队列

	// 将入度为 0 的课程加入队列
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	// 拓扑排序过程
	for len(q) > 0 {
		u := q[0] // 取出队列中的课程
		q = q[1:] // 队列出队
		result = append(result, u) // 将课程加入拓扑排序结果中

		// 遍历当前课程的所有后续课程
		for _, v := range edges[u] {
			indeg[v]-- // 将后续课程的入度减少
			if indeg[v] == 0 {
				q = append(q, v) // 如果入度为 0，将后续课程加入队列
			}
		}
	}

	// 如果拓扑排序结果包含所有课程，则可以完成课程学习
	return len(result) == numCourses
}
