package main


import "fmt"

func main(){
	fmt.Println(jump([]int{2,3,1,1,4}))
}

func jump(nums []int) int {
	end := 0
	position := 0
	steps := 0
	for i := 0; i < len(nums) - 1; i++{
		position = Max(position, nums[i] + i)
		if i == end{
			end = position
			steps++
		}
	}
	return steps
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}