package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	//排序
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	//map
	fmt.Println(isAnagram1("anagram", "nagaram"))
	fmt.Println(isAnagram1("rat", "car"))
	//数组
	fmt.Println(isAnagram2("anagram", "nagaram"))
	fmt.Println(isAnagram2("rat", "car"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sArr := strings.Split(s,"")
	tArr := strings.Split(t,"")
	sort.Strings(sArr)
	sort.Strings(tArr)
	if reflect.DeepEqual(sArr, tArr)  {
		return true
	}
	return false
}


func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letterCount := make(map[uint8]int)
	for i := 0; i < len(s);  i++ {
		letterCount[s[i]] = letterCount[s[i]] + 1
	}
	for i := 0; i < len(t);  i++ {
		letterCount[t[i]] = letterCount[t[i]] - 1
	}
	flag := true
	for _ ,v := range letterCount{
		if v != 0 {
			flag = false
			break
		}
	}
	return flag
}


func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letterArrS := [26]int{}
	letterArrT := [26]int{}
	for i := 0; i < len(s);  i++ {
		key := s[i] - 'a'
		letterArrS[key] = letterArrS[key] + 1
	}
	for i := 0; i < len(t);  i++ {
		key := t[i] - 'a'
		letterArrT[key] = letterArrT[key] + 1
	}
	if letterArrS == letterArrT  {
		return true
	}
	return false
}

