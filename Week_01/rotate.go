package Week_01

import "fmt"

func main() {
	rotate([]int{1,2,3,4,5,6,7},3)
	rotate1([]int{1,2,3,4,5,6,7},3)
}

func rotate1(nums []int, k int)  {ss
	if k > 0 && len(nums) > 1 {
		k = k % len(nums)
		nums_new := append(nums[len(nums) - k:],nums[0:len(nums) - k]...)
		nums = append(nums[:0], nums_new...)
		fmt.Println(nums)
	}
}

func rotate(nums []int, k int)  {
	lenth := len(nums)
	if k > 0 && lenth > 1 {
		k = k % lenth
		nums_new := []int{}
		for i := lenth-k ; i <= lenth-1 ; i++ {
			nums_new = append(nums_new, nums[i])
		}
		for i := 0 ; i < lenth-k ; i++ {
			nums_new = append(nums_new, nums[i])
		}
		copy(nums,nums_new)
		fmt.Println(nums)
	}
}