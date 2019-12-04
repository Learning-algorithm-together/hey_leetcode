package main

import (
	"fmt"
	"math"
)

func main() {
	sli := []int{1,8,6,2,5,4,8,3,7}
	maxArea(sli)
}
// 暴力解法，循环遍历，算出最大值，算法复杂度O(N^2)。
// 由第一开始，循环遍历减去后面每个值，求面积，保留最大值。
// 然后从第二个开始，继续往后遍历，保留最大值。
/*
func maxArea(height []int) int {
	// 最大值
	max := 0.0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			// 求出面积
			height := (height[i] - height[j])
			area := height * height
			fmt.Println("i = ",i , "j = ", j, "height = ", height, "area = ", area)
			max = math.Max(max, float64(area))
		}
	}
	fmt.Println("max = ",max)
	return int(max)
}
*/
// 左右边界，向中间收敛，矮的移动，高的不动
// 夹逼法
/*
func maxArea(height []int) int {
	// 最大值
	max := 0.0
	// 初始化j
	j := len(height) - 1
	i := 0
	for {
		// 一旦i和j相遇，跳出循环
		if i >= j {
			break
		}
		if height[i] < height[j] {
			area := float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
			max = math.Max(max, area)
			i++
		}else {
			area := float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
			max = math.Max(max, area)
			j--
		}
	}
	
	return int(max)
}
*/

func maxArea(height []int) int {
	// 最大值
	max := 0.0
	// 初始化j
	for i, j := 0, (len(height) - 1); i < j;  {
		fmt.Println("i = ", i, "; j = ", j)
		var min int
		if height[i] < height[j] {
			fmt.Println("height[i+1] = ", height[i+1])
			min = height[i]
			i++
		}else {
			fmt.Println("height[j-1] = ", height[j-1])
			min = height[j]
			j--
		}
		area := (j - i +1)*min
		fmt.Println("i = ", i, "j = ", j, "area = ",area)
		max = math.Max(float64(max), float64(area))
		fmt.Println("max = ",max)
	}
	return int(max)
}
