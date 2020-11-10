package datastruct

import "fmt"

/**
1. 冒泡排序
就像泡泡一样，每次循环一次，比较大于的话就交换位置，最大数会到最上面。
时间复杂度：n方
空间复杂度：没有额外空间
稳定性(两个相同的数，不交换位置就是稳定的)：稳定
*/
func bubbleDemo(s []int) {
	length := len(s)
	var flag bool //表示是否有数据交换
	for i := 0; i < length; i++ {
		flag = false
		for j := 0; j < length-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(s)
}

/**
2.1 插入排序
一半有序，一半无序，将无序的插进有序中。
1，3，5，8，2，1
加入对将2存在一个临时变量temp中， 然后指针从2的前一位开始判断，如果前一位比2大，则将前一位往后移，
*/
func insertDemo(s []int) {
	length := len(s)

	for i := 1; i < length; i++ {
		value := s[i]
		j := i - 1 //j是有序的一半，每次循环后,排好序的就多了
		for ; j >= 0; j-- {
			if value < s[j] { //小于拍好的序的数，交换位置
				s[j+1] = s[j] //这里没交换位置，只是将前面的数覆盖了后面的。这样写是为了不使用中间变量。提高速度。
			} else {
				break
			}
		}
		s[j+1] = value //这里将值放到最终位置
	}
	fmt.Println(s)
}

func insertDemo000(s []int) {
	length := len(s)
	for i := 0; i < length; i++ {

	}
}

//2.2 写法2
func insertDemo2(s []int) {
	length := len(s)

	for i := 1; i < length; i++ {
		j := i - 1    //j是有序的一半，每次循环后,排好序的就多了
		value := s[i] //这里必须把值拿出来，因为a[i]的值在交换的过程中可能变化，但是我们要求他是不能改变的。
		for ; j >= 0; j-- {
			if value < s[j] {
				s[j], s[j+1] = s[j+1], s[j] //go里面提供了这样的语句，可以直接一步到位，到底性能怎么样，我们就不得而知了。
			} else {
				break
			}
		}
	}
	fmt.Println(s)
}

/**
3. 选择排序
选择排序每次都要找剩余未排序元素中的最小值，并和前面的元素交换位置，这样破坏了稳定性。
选择排序是不稳定的。相等的数，排序以后，前后关系位置会变化。
*/

func selectDemo(s []int) {
	length := len(s)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i] //或者将中间变量保存起来最后在交换。
			}
		}
	}
	fmt.Println(s)
}

/**
总结：
冒泡：每次循环都比较一次，从最下面，和每个数都比较，选出最大的数继续和下一个数比较，直到数组最后。
插入：两个序列，排序好的，未排序好的，每次，从未排序好的，拿出来数跟排序好的中的数从后向前做比较，插到比他小的后面。
选择：两个序列，排序好的，为排序好的，每次，从为排序好的中选出最小值，然后放到排序好的后面。

稳定的排序：冒泡，插入
平时的选择：插入，相对来说比冒泡更少的交换。
*/
