package main

import "fmt"

func main (){
	//题目转化：有k个各自可以放1~n中任意数字，有多少种方法
	fmt.Println(combine(4,2))
}

func combine(n int, k int) [][]int {
	answer := [][]int{}
	result :=[]int{}
	combineArr(n,k,result,&answer)
	return answer
}
func combineArr(n int, k int, result []int, answer *[][]int) {
	var length = len(result)
	if length >= k {
		temp := make([]int,k)
		copy(temp,result)
		*answer = append(*answer,temp)
		return
	}
	var i int
	if length == 0{
		i = 1
	}else{
		i = result[length-1]+1
	}
	for i <= n {
		combineArr(n,k,append(result,i),answer)
		i++
	}

}
