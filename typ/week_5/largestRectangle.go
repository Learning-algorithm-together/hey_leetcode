package main

import (
	"fmt"
	"math"
)

func main() {
	sli := []int{999,999,999,999}
	fmt.Println(largestRectangleArea(sli))
}

// 暴力法
func largestRectangleArea(heights []int) int {
	var lar int
	for i := 0; i < len(heights); i++ {
		minHeiht := math.MaxInt64
		for j := i; j < len(heights); j++ {
			var maxArea int
			if heights[j] < minHeiht {
				minHeiht = heights[j]
				fmt.Println("heights[j] = ")
			}
			maxArea = (j-i+1) * minHeiht
			fmt.Println("i < j: i = ", i, ", j = ", j, "maxArea = ", maxArea, ", minHeiht = ", minHeiht)
			if lar < maxArea {
				lar = maxArea
			}
		}
	}
	return lar
}