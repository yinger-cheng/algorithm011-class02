package Week_01

import "fmt"

//func main() {
//	fmt.Println(twoSum([]int{3,2,4},6))
//}

func twoSum(nums []int, target int) []int {
	result := []int{}
	length := len(nums)
	if length < 2 {
		return result
	}
	for i := 0; i < length -1; i++{
		j := i+1
		for j < length {
			fmt.Println(nums[i] + nums[j])
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}else{
				j++
			}
		}

	}

	return result
}
