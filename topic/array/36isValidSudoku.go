package main

/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。


示例 1：


输入：board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：true
示例 2：

输入：board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
输出：false
解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。


提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'
解题思路
对应每一行、每一列、每个三宫格 都创建一个map
key=number由于number=>1-9,用数组来代替map

其中

rows[i][number] = times 表示i行某个number出现了times次

columns[j][number] = times 表示j列某个number出现了times次

subboxs[i/3][j/3][number] = times 一共9个小宫格,第i/3行第j/3列个宫格 中 的hashMpa中 number出现了times次

遍历整个数组，遇到一个数字，分别将他对应的行HashMap,列hashMap，宫格hashMap添加一个数字

代码
由于原数组是byte类型的数组和空格，因此

遇到空格跳过不操作

遇到数字将byte类型 => int类型（byte存储的字符集中的编号）
func isValidSudoku(board [][]byte) bool {
    // 1. 数字1-9， 第二个位置表示number：10 => 0-9
    rows, columns := [9][10]int{}, [9][10]int{}
    subboxs := [3][3][10]int{}
    // 2. 遍历数组
    for i, row := range board {
        for j, num := range row {
            // 空格不用管
            if num == '.' {
                continue
            }
            number := num - '1' //将byte换算成数值

            // 判断某元素是否出现过>1次
            if (rows[i][number] > 0) || (columns[j][number] > 0) || (subboxs[i/3][j/3][number] > 0) {
                return false
            }
            // 没有出现过，则找打对应的行列宫格map中对应的number次数++
            rows[i][number]++
            columns[j][number]++
            subboxs[i/3][j/3][number]++
        }
    }
    return true
}


作者：BigRainKing
链接：https://leetcode.cn/problems/valid-sudoku/solutions/1600358/by-lin-xia-5-6p9i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9 ; i++ {
		x_result := make(map[byte]bool , 9)
		y_result := make(map[byte]bool , 9)
		box_result := make(map[byte]bool , 9)

		for j :=0; j<9; j++ {
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if board[i][j] != '.'  {
				if x_result[board[i][j]] {
					return false
				}
			}

			if board[row][col] != '.' {
				if box_result[board[row][col]] {
return false
				}
			}
			if board[j][i] != '.' {
				if y_result[board[j][i]] {
					return false
				}
			}
			box_result[board[row][col]] = true
			x_result[board[i][j]] = true
			y_result[board[j][i]] = true
		}

	}
	return true
}

// 解法二 添加缓存，时间复杂度 O(n^2)
func isValidSudoku1(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0' - byte(1)
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
}

func main(){
  i,j:= 8,8
	row := (i%3)*3 + j%3
	col := (i/3)*3 + j/3
	println(row)
	println(col)
	//println(i%3)
}