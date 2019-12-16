package jiangjin

import (
	"testing"
)

func TestISVaild(t *testing.T){

}


type MyStack struct {
	flag int
	data []Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		flag:0,
		data:[]Queue{
			NewQueue(),
			NewQueue(),
		},
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.data[this.flag].Push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.data[this.flag].Len() == 1{
		return this.data[this.flag].Pop()
	}
	for this.data[this.flag].Len() != 1{
		item := this.data[this.flag].Pop()
		if this.flag == 1{
			this.data[0].Push(item)
		}else{
			this.data[1].Push(item)
		}
	}

	// 切换队列
	item := this.data[this.flag].Pop()
	if this.flag == 1{
		this.flag = 0
	}else{
		this.flag = 1
	}
	return item
}


/** Get the top element. */
func (this *MyStack) Top() int {
	x := this.Pop()
	this.Push(x)
	return x
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data[this.flag].data) == 0
}

type Queue struct{
	data []int
}

func NewQueue() Queue{
	return Queue{data:[]int{}}
}

func (q *Queue) Push(x int){
	q.data = append(q.data, x)
}

func (q *Queue) Pop() int {
	if q.IsEmpty(){
		return 0
	}
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *Queue) Peek() int {
	return q.data[0]
}

func (q *Queue) Len() int{
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}