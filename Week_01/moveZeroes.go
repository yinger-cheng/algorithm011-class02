package Week_01

import (
	"fmt"
)

func main() {
	nums := []int{0,2,3,6,0,1,0,8}
	moveZeroes(nums)
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
		fmt.Println("------------")
		fmt.Println(i,nums)
	}
	fmt.Println(nums)
	return
}

