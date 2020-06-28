package Week_01

import "fmt"


func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,3,3,4,4,4,5,5,6,6,7,7,7,7,8}))
	fmt.Println(removeDuplicates1([]int{1,1,1,2,2,3,3,4,4,4,5,5,6,6,7,7,7,7,8}))
}

func removeDuplicates(nums []int) int {
	for i := 0 ; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func removeDuplicates1(nums []int) int {
	for i := 0 ; i < len(nums)-1 ; i++ {
		j := i+1
		//下表为1的值和前一值
		if nums[i] == nums[j]{
			for nums[j] == nums[i] && j < len(nums) {
				if j == len(nums)-1{ //比较到最后一个值时不再++
					break
				}
				j++
			}
			if nums[i] == nums[j]{
				nums = nums[:i+1]
			}else{
				nums = append(nums[:i+1], nums[j:]...)
			}
		}
	}
	return len(nums)
}


