package main

import "fmt"

/**
* Description:
* @Author Hollis
* @Create 2023-10-14 22:24
 */

func solveMaze(maze [][]int, startX, startY, endX, endY int) bool {
	if startX == endX && startY == endY {
		// 已经到达终点
		return true
	}

	if isValidMove(maze, startX, startY) {
		// 标记当前位置为已访问
		maze[startX][startY] = 2

		// 尝试向上、向下、向左、向右四个方向移动
		if solveMaze(maze, startX+1, startY, endX, endY) || // 向下
			solveMaze(maze, startX-1, startY, endX, endY) || // 向上
			solveMaze(maze, startX, startY+1, endX, endY) || // 向右
			solveMaze(maze, startX, startY-1, endX, endY) { // 向左
			return true
		}

		// 如果上述四个方向都没有通路，则回退
		maze[startX][startY] = 0
	}

	return false
}

func isValidMove(maze [][]int, x, y int) bool {
	// 检查是否越界或者是否是墙（1 表示墙，0 表示通路，2 表示已访问）
	if x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == 0 {
		return true
	}
	return false
}

func main() {
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	startX, startY := 0, 0
	endX, endY := len(maze)-1, len(maze[0])-1

	if solveMaze(maze, startX, startY, endX, endY) {
		fmt.Println("找到一条路径：")
		printMaze(maze)
	} else {
		fmt.Println("未找到路径。")
	}
}

func printMaze(maze [][]int) {
	for _, row := range maze {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print("  ") // 通路
			} else if cell == 1 {
				fmt.Print("█ ") // 墙
			} else if cell == 2 {
				fmt.Print("x ") // 已访问
			}
		}
		fmt.Println()
	}
}
