package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2},0))
}
func fourSum(nums []int,target int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	result := [][]int{}
	keyMap := make(map[string]int)
	if length < 4 || nums == nil {
		return result
	}
	for i := 0;i < length -3  ; i++  {
		if nums[i] > target {
			continue
		}
		j := i + 1
		k := length - 1
		for j < k {
			fmt.Println("------")
			fmt.Println(nums[i] + nums[j] + nums[k])
			sum := nums[i] + nums[j] + nums[k]
			if sum < target {
				j++
			}else if sum > target {
				k--
			}else{
				//keyStr := fmt.Sprintf("%v_%v_%v",nums[i] , nums[j],nums[k])
				keyStr := strconv.FormatInt(int64(nums[i]), 10) + strconv.FormatInt(int64(nums[j]), 10)+strconv.FormatInt(int64(nums[k]), 10)
				fmt.Println(keyStr)
				if keyMap[keyStr] != 1 {
					result = append( result,[]int{ nums[i] , nums[j], nums[k]} )
					keyMap[keyStr] = 1
				}
				j++
				k--
			}
		}
	}
	return result
}

