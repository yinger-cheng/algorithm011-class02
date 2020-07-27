package main

import "fmt"

func main(){
	fmt.Println(lemonadeChange([]int{5,5,5,10,5,5,10,20,20,20}))
}

func lemonadeChange(bills []int) bool {
	five ,ten := 0 , 0
	for _, v := range bills {
		if v == 5 {
			five ++
		}else if v == 10 {
			if five < 1 {
				return false
			}
			five --
			ten ++
		}else if v == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			}else {
				return false
			}
		}
	}
	return true
}
