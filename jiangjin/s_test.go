package jiangjin

import (
	"fmt"
	"sort"
	"testing"
)

func TestISVaild(t *testing.T){
	fmt.Println(isAnagram("a", "b"))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	result := make([][]string, 0)
	for i:=0; i<len(strs); i++{
		s := SortStringByCharacter(strs[i])
		m[s] = append(m[s], strs[i])
	}

	for _,v:=range m{
		result = append(result, v)
	}
	return result
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
