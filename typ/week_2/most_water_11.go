package main

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
/*
func maxArea(height []int) int {
	// 最大值
	max := 0
	// 初始化j
	for i, j := 0, (len(height) - 1); i < j;  {
		tmp := 0
		if height[i] < height[j] {
			tmp = (j - i)*height[i]
			i++
		}else {
			tmp = (j - i)*height[j]
			j--
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
*/

func maxArea(height []int) int {
	// 最大值
	max := 0
	i := 0
	j := len(height) - 1
	// 初始化j
	for  i != j {
		tmp := 0
		if height[i] < height[j] {
			tmp = (j - i)*height[i]
			i++
		}else {
			tmp = (j - i)*height[j]
			j--
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// 用时4ms的案例
/*
func maxArea(height []int) int {
	head := 0
	tail := len(height) - 1
	max := 0
	for head != tail {
		tmp := 0
		if height[head] < height[tail] {
			tmp = (tail - head) * height[head]
			head++
		} else {
			tmp = (tail - head) * height[tail]
			tail--
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
*/
