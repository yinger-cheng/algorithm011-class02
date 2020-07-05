package main

import (
	"fmt"
)

func main() {
	nums := []int{0,2,3,6,0,1,0,8}
	moveZeroes(nums)
	moveZeroes1(nums)
}


func moveZeroes(nums []int) {
	index := 0
	for i := 0 ;i < len(nums) -1 ;i++{
		if nums[index] == 0 {
			nums =  append(nums[:index], nums[index+1:]...)
			nums = append(nums,0)
		}else{
			index ++
		}
	}
	fmt.Println(nums)
	return
}

func moveZeroes1(nums []int) {
	j := 0
	for i := 0 ;i < len(nums) ;i++{
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
	fmt.Println(nums)
	return
}

