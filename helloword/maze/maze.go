package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

// var dirs [4]point = [4]point [4]poin和[]point是不同类型
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 判断是否在坐标范围内
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 不在坐标范围内或者遇到墙
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}

			// 判断是否是走过的点，防止走回去
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}

			// 如果返回原点
			if next == start {
				continue
			}

			// 记录步数，然后往队列添加符合条件的点
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			Q = append(Q, next)

		}
	}

	return steps
}

func main() {
	maze := readMaze("./maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
