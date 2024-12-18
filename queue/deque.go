package queue

import "gostl/list"

// 双端队列，基于双向链表的实现
type Deque[T any] struct {
	list list.DList[T]
}

// NewDeque 构造一个空的双端队列
func NewDeque[T any]() Deque[T] {
	return Deque[T]{}
}

// Len 获取双端队列长度
func (q Deque[T]) Len() int {
	return q.list.Len()
}

// Empty 判断双端队列是否为空
func (q Deque[T]) Empty() bool {
	return q.list.Empty()
}

// Clear 清空双端队列
func (q *Deque[T]) Clear() {
	q.list.Clear()
}

// Front 获取双端队列队首元素
func (q Deque[T]) Front() T {
	return q.list.Front()
}

// Back 获取双端队列队尾元素
func (q Deque[T]) Back() T {
	return q.list.Back()
}

// PushFront 在双端队列队首插入元素
func (q *Deque[T]) PushFront(value T) {
	q.list.PushFront(value)
}

// PushBack 在双端队列队尾插入元素
func (q *Deque[T]) PushBack(value T) {
	q.list.PushBack(value)
}

// PopFront 删除并返回双端队列队首元素
func (q *Deque[T]) PopFront() T {
	return q.list.PopFront()
}

// PopBack 删除并返回双端队列队尾元素
func (q *Deque[T]) PopBack() T {
	return q.list.PopBack()
}
