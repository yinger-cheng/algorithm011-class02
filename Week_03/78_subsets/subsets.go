package main

import "fmt"

func main(){
	fmt.Println(subsets([]int{1,2,3}))
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	for _, v := range nums {
		temp := make([][]int, len(result))
		for key, v1 := range result {
			v1 = append(v1, v)
			temp[key] = append(temp[key], v1...)
			result = append(result, temp[key])
		}
	}
	return result
}