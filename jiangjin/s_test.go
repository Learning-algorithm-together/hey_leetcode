package jiangjin

import (
	"fmt"
	"sort"
	"testing"
)

func TestISVaild(t *testing.T){
	fmt.Println(isAnagram("a", "b"))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	m := make(map[uint8]int)
	for i:=0;i<len(s);i++{
		m[s[i]]++
		m[t[i]]--
	}

	for _, v:= range m{
		if v != 0{
			return false
		}
	}

	return true
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

