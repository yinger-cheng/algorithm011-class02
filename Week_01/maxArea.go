package Week_01

//func main() {
//	nums := []int{1,8,6,2,5,4,8,3,7}
//	fmt.Println(maxArea(nums))
//	fmt.Println(maxArea1(nums))
//}

func maxArea(height []int) int {
	var maxArea int
	for i :=0; i < len(height)-1; i++  {
		for j := len(height)-1 ; j > i; j-- {
			maxArea = Max(maxArea,(j-i) * Min(height[i],height[j]))
		}
	}
	return maxArea
}

func maxArea1(height []int) int {
	var maxArea int
	i := 0
	j := len(height)-1
	for  ; i < j;  {
		maxArea = Max(maxArea,(j-i) * Min(height[i],height[j]))
		if height[i] <= height[j] {
			i++
		}else{
			j--
		}
	}
	return maxArea
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}