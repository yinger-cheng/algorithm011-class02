package main

import "fmt"

func main(){
	fmt.Println(canJump([]int{2,3,1,1,5}))
}

func canJump(nums []int) bool {
	length := len(nums) - 1
	for i:= length - 1;i >= 0; i-- {
		if nums[i] + i >= length {
			length = i
		}
	}
	return length <= 0
}
