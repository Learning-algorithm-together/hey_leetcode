package main

import "fmt"

func main() {
	climbStairs(20)
	fmt.Println(climbStairs(20))
}
// 采用递归的思想
/*
	1:1
	2:2
	3:f(1)+f(2)=3
	4:f(3)+f(2)=5
	n:f(n-1)+f(n-2)
*/
// 超级慢，简直要完球
/*/
func climbStairs(n int) int {

	i := climb_stairs(n)
	return i
}
var count int
func climb_stairs(i int) int {

	if i <= 1 {
		return 1
	}
	fmt.Println("climb i = ", count)
	count++
	return climb_stairs(i-1) + climb_stairs(i-2)
}
*/
// 存储中间变量，减少计算次数
func climbStairs(n int) int {
	var slic = make([]int, n)
	i := climb_stairs(n, slic)
	return i
}
var count int
func climb_stairs(i int, slic []int) int {

	if i <= 1 {
		return 1
	}
	if i == 2 {
		return 2
	}
	if slic[i-1] == 0 {
		slic[i-1] = climb_stairs(i-1, slic) + climb_stairs(i-2, slic)
	}
	fmt.Println("count = ", count)
	count++
	return slic[i-1]
}
