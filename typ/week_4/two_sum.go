package main

import "fmt"

func main()  {
	nums := []int{1, 3, -4, 7, -3, 0}
	fmt.Println(twoSum(nums, 3))
}
// 暴力法
/*
	1.循环遍历
	2.双重循环，算法复杂度N^2
*/
/*
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}
*/
// 两遍哈希表
/*
	1.第一遍遍历，将所有数据映射到hash map
	2.第二遍遍历，将所需的数字从map中取得，有就返回，没有就下一个
 */
/*
func twoSum(nums []int, target int) []int {
	temp := make(map[int]interface{}, len(nums))
	for index, value := range nums {
		temp[value] = index
	}
	for i := 0; i < len(nums); i++ {
		v := target - nums[i]
		if _, ok := temp[v];ok && i != temp[v].(int){
			return []int{i,temp[v].(int)}
		}
	}
	return nil
}
*/
// 一遍哈希表
/*
	1.遍历的同时插入map
	2.向map中寻找符合条件的数
 */

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int,len(nums))
	for index, value := range nums {
		i := target - value
		if _, ok := temp[i]; ok {
			return []int{temp[i], index}
		}else {
			temp[value] = index
		}
	}
	return nil
}