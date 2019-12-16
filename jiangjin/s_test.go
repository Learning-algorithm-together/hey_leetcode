package jiangjin

import (
	"fmt"
	"testing"
)

func TestISVaild(t *testing.T){
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	//"getMin","pop","top","getMin"
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
}
type MinStack struct {
	data []int
	// 依次放入更小的
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min: 1<<63 - 1,
	}
}

func (this *MinStack) Push(x int)  {
	if x <= this.min{
		this.data = append(this.data, x)
		this.min = x
	}
	this.data = append(this.data, x)
}

func (this *MinStack) Pop()  {
	if this.min == this.data[len(this.data)-1]{
		this.data = this.data[:len(this.data)-1]
		this.min =  this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
