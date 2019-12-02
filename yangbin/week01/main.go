
/*
	例如有8支球队，两两组合的选择有多少种？ (8-1)*(8-3)*(8-5)*(8-7)=105
*/

package main

import (
	"fmt"
)

func Choose(n int) (res int) {
	res = 1
	for i:=1; i < n; i++ {
		if i % 2 != 0 {
			res *= (n-i)
		}
	}
	return
}

func main() {
	n := 8
	if n % 2 != 0 {
		res := "n 不是偶数，不符合规范。"
		fmt.Println(res)
		return
	}

	fmt.Println(Choose(n))
}