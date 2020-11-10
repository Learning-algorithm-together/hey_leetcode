package leetcode

/**

如何基于链表实现 LRU 缓存淘汰算法？
我的思路是这样的：我们维护一个有序单链表，越靠近链表尾部的结点是越早之前访问的。
当有一个新的数据被访问时，我们从链表头开始顺序遍历链表。
1. 如果此数据之前已经被缓存在链表中了，我们遍历得到这个数据对应的结点，并将其从原来的位置删除，然后再插入到链表的头部。
2. 如果此数据没有在缓存链表中，又可以分为两种情况：如果此时缓存未满，
则将此结点直接插入到链表的头部；如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部。

*/

type Object interface{}

//type Node struct {
//	data Object
//	next *Node
//	list *List
//}
//
//type List struct {
//	size uint64
//	head *Node
//	tail *Node
//}
