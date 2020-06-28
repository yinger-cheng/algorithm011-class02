package Week_01

//func main() {
//fmt.Println(climbStairs(7))
//}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	//f1 := 1
	f2 := 2
	f3 := 3
	total := 0
	for i := 4 ;i <= n; i++ {
		total = f2 + f3
		f2 = f3
		f3 = total
	}
	return total

}