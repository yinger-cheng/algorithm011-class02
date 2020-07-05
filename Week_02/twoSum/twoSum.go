package main

import "fmt"

func main(){
	fmt.Println(twoSum([]int{2, 7, 11, 15},9))
	fmt.Println(twoSum([]int{3,2,4},6))
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	numsMap := make(map[int]int)
	for k ,v := range nums{
		numsMap[v] = k
	}
	for i := 0 ; i < len(nums) ; i++ {
		searchNum := target - nums[i]
		if numsMap[searchNum] != 0 && i != numsMap[searchNum] {
			return []int{i,numsMap[searchNum]}
		}
	}
	return result
}