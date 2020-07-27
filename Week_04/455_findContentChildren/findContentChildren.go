package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(findContentChildren([]int{1,2,3,5},[]int{3}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for _, v := range g {
		for j, v2 := range s {
			if v <= v2 {
				count = count + 1
				s = s[j+1:]
				break
			}
		}
	}
	return count
}