package Week_01

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{2,4,9,3,9}))
}

func plusOne(digits []int) []int {
	j := 0
	lastNum := digits[len(digits)-1]
	lastIndex := len(digits) -1
	if lastNum == 9 {
		digits[lastIndex] = 0
		j = 1
	}else{
		digits[lastIndex] = digits[lastIndex] + 1
	}
	if j > 0 {
		for i := len(digits)-2 ; i >= 0 ; i-- {
			if digits[i] == 9  && j == 1{
				digits[i] = 0
				j = 1
			}else{
				digits[i] = digits[i] + 1
				j = 0
				break
			}
		}
		if j == 1 {
			digits = append([]int{1},digits...)
		}
	}
	return digits
}
