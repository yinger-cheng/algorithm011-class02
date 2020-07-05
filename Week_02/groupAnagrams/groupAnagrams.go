package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	string := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(string))
}

func groupAnagrams(strs []string) [][]string {
	wordsMap := make(map[string][]string)
	for i := 0 ;i < len(strs)  ; i++ {
		words := strings.Split(strs[i],"")
		sort.Strings(words)
		keyword := strings.Join(words, ``)
		wordsMap[keyword] = append(wordsMap[keyword],strs[i])
	}
	result := [][]string{}
	for _,v := range wordsMap{
		result = append(result,v)
	}
	return result
}
