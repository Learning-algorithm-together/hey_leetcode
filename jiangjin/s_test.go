package jiangjin

import (
	"fmt"
	"testing"
)

func TestISVaild(t *testing.T){
	fmt.Println(isValid("){"))
}

func isValid(s string) bool {
	if len(s) == 1{
		return false
	}
	if s == ""{
		return true
	}
	// 申请一个栈
	stack := make([]uint8, len(s))
	stackLen := 0
	for i:=0; i<len(s); i++{

		// 左括号入栈
		if s[i] == '(' || s[i] == '{' || s[i] == '['{
			stack[stackLen] = s[i]
			stackLen++
		}else{
			if stackLen == 0{
				return false
			}


			//  判断是否成对
			if s[i] == ')' && stack[stackLen-1] == '(' {
				stackLen--
			} else if s[i] == '}' && stack[stackLen-1] == '{' {
				stackLen--
			} else if s[i] == ']' && stack[stackLen-1] == '['{
				stackLen--
			}else{
				return false
			}
		}
	}

	return  stackLen == 0
}
