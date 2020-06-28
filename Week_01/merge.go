package Week_01

import "sort"

func main() {
	merge([]int{1,2,3,0,0,0,0,0,0,0},3,[]int{0,1,2,5,6},5)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

