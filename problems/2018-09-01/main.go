package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

func solution(maze [][]int, start []int, finish []int) int {
	maze[start[0]][start[1]] = -1
	mark(maze, start, 0)
	printMaze(maze)
	return maze[finish[0]][finish[1]]
}

type pos struct {
	i int
	j int
}

// marks unmarked point in maze with n, and all neibors n+1
// return true if point was marked
func mark(maze [][]int, point []int, n int) {
	i, j := point[0], point[1]

	neibors := map[pos]bool{
		pos{i + 1, j}: false,
		pos{i - 1, j}: false,
		pos{i, j - 1}: false,
		pos{i, j + 1}: false,
	}
	for p := range neibors {
		neibors[p] = markP(maze, p.i, p.j, n+1)
	}

	for p, ok := range neibors {
		if ok {
			mark(maze, []int{p.i, p.j}, n+1)
		}
	}
}

// mark maze[i][j] if exists and not marked
func markP(maze [][]int, i, j, n int) bool {
	if i >= len(maze) || j >= len(maze[0]) {
		return false
	}
	if i < 0 || j < 0 {
		return false
	}
	if maze[i][j] != 0 {
		return false
	}

	maze[i][j] = n
	return true
}

func printMaze(maze [][]int) {
	for i, row := range maze {
		for j := range row {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}
}
