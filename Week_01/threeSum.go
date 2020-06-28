package Week_01

import (
	"fmt"
	"sort"
	"strconv"
)

func main (){
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0,0,0,0}))
	fmt.Println(threeSum1([]int{-2,0,1,1,2}))
}

func threeSum(nums []int) [][]int {
	lenth := len(nums)
	result := [][]int{}
	fmt.Println(lenth)
	keyMap := make(map[string]int)
	if lenth < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < lenth -2;i++ {
		for j := i+1; j < lenth -1 ;j++ {
			for k := j+1; k < lenth  ;k++ {
				keyStr := fmt.Sprintf("%v_%v_%v",nums[i] , nums[j] , nums[k])
				if nums[i] + nums[j] + nums[k] == 0 && keyMap[keyStr] != 1{
					keyMap[keyStr] = 1
					result = append( result,[]int{nums[i] , nums[j] , nums[k]} )
				}
			}
		}
	}
	return result
}

func threeSum1(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	result := [][]int{}
	keyMap := make(map[string]int)
	if length < 3 || nums == nil {
		return result
	}
	for i := 0;i < length -2  ; i++  {
		if nums[i] > 0 {
			break
		}

		j := i + 1
		k := length - 1
		for j < k {
			fmt.Println("------")
			fmt.Println(nums[i] + nums[j] + nums[k])
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			}else if sum > 0 {
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

