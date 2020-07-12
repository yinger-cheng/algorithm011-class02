package main

import "fmt"

func main (){
	fmt.Println(generateParenthesis(1))
}
func generateParenthesis(n int) []string {
	result := generate(0,0,  n , "",[]string{})
	return result
}
//left 随时可以加 只要别超标
// right 左个数 > 右个数
func generate(left int,right int ,n int ,s string,result []string) []string{
	if left == n && right == n {
		result = append(result,s)
		return result
	}
	if left < n {
		result = generate(left + 1 ,right, n, s + "(",result)
	}
	if left > right {
		result = generate(left, right + 1, n, s + ")",result)
	}
	return result
}