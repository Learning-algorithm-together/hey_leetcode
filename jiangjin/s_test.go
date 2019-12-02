package jiangjin

import (
	"testing"
)

// 测试代码

func Test_Test(t *testing.T){
	moveZeroes([]int{0,1,0,3,12})
}



func moveZeroesV2(nums []int){
	for i:=0;i<len(nums);i++{
		if nums[i] != 0{
			continue
		}
		// 如果是0需要向后查找不是0的元素替换
		for j:=i+1;j<len(nums);j++{
			if nums[j] == 0{
				continue
			}else{
				nums[i],nums[j] = nums[j],nums[i]
				break
			}
		}

	}
}

func moveZeroesV3(nums []int)  {
	// 使用两个指针
	// i ,j 两个一起走
	// 当j遇到0 时就不懂,i 继续走 遇到不为0的元素和j 进行交换
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[i] != 0{
			nums[i],nums[j] = nums[j],nums[i]
			j++
		}
	}
}

func moveZeroes(nums []int)  {
	// 使用两个指针
	// i ,j 两个一起走
	// 当j遇到0 时就不懂,i 继续走 遇到不为0的元素和j 进行交换
	last := 0
	for i:=0;i<len(nums);i++{
		if nums[i] != 0{
			nums[last] = nums[i]
			last++
		}
	}

	for {
		if last <len(nums){
			nums[last] = 0
			last++
		}else{
			break
		}
	}
}