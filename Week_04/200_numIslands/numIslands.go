package main

import "fmt"

func main(){
	input := [][]byte{
		[]byte{'1','1','1','1','0'},
		[]byte{'1','1','0','1','0'},
		[]byte{'1','1','0','0','0'},
		[]byte{'0','0','0','0','0'},
		}
	fmt.Println(numIslands(input))
	input2 := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(input2))
}


func numIslands(grid [][]byte) int {
	count := 0
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	for i :=0 ; i < n; i++ {
		for j := 0 ; j < m ; j++ {
			if grid[i][j] == '1' {
				DFSMarking(grid,i,j,n,m)
				count++
			}
		}
	}
	return count
}

func DFSMarking(grid [][]byte, i int, j int ,n int, m int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	DFSMarking(grid,i+1,j,n,m)
	DFSMarking(grid,i-1,j,n,m)
	DFSMarking(grid,i,j+1,n,m)
	DFSMarking(grid,i,j-1,n,m)
}