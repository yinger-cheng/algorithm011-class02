package main

import "fmt"

func main(){
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	//result := []string{}
	result := new([]string)
	generate(0,0, n,"", result)
	return *result
}

func generate(left int, right int,max int ,str string, result *[]string) {
	if left == right && right == max {
		*result = append(*result,str)
	}
	if left < max {
		generate(left + 1, right, max, str + "(", result)
	}
	if right < left {
		generate(left, right + 1, max, str + ")", result)
	}
}