package main

import "fmt"

func main() {
	start := "hit"
	end := "cog"
	bank := []string{"hot","dot","dog","lot","log","cog"}
	fmt.Println(ladderLength(start, end, bank))
}
func idxOf(str string, bank []string) int{
	for i , s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord,wordList) == -1 {
		return 0
	}


	return 0
}
