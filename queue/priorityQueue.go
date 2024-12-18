package queue

import (
	"gostl"
	"gostl/heap"
)

type PriorityQueue[T any] struct {
	heap []T
	impl pqImpl[T]
}

func NewPriorityQueue[T gostl.Ordered]() *PriorityQueue[T] {
	pq := pqOrdered[T]{}
	pq.impl = (pqImpl[T])(&pq)
	return &pq.PriorityQueue
}

func NewPriorityQueueInitializerList[T gostl.Ordered](values ...T) *PriorityQueue[T] {
	heap.NewMinHeap(&values)
	pq := pqOrdered[T]{}
	pq.heap = values
	pq.impl = (pqImpl[T])(&pq)
	return &pq.PriorityQueue
}

func NewPriorityQueueFunc[T any](less gostl.LessFunc[T]) *PriorityQueue[T] {
	pq := pqFunc[T]{}
	pq.less = less
	pq.impl = (pqImpl[T])(&pq)
	return &pq.PriorityQueue
}

func NewPriorityQueueFuncInitializerList[T any](less gostl.LessFunc[T], values ...T) *PriorityQueue[T] {
	heap.NewMinHeapFunc(&values, less)
	pq := pqFunc[T]{}
	pq.less = less
	pq.heap = values
	pq.impl = (pqImpl[T])(&pq)
	return &pq.PriorityQueue
}

// Len 获取当前 priority queue 节点数
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

// Empty 获取 priority queue 是否为空
func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.heap) == 0
}

// Clear 清空当前 priority queue
func (pq *PriorityQueue[T]) Clear() {
	pq.heap = pq.heap[:0]
}

// Top 获取 priority queue 头部元素，若 priority queue 为空则 panic
func (pq *PriorityQueue[T]) Top() T {
	return pq.heap[0]
}

// Push 向 priority queue 插入元素
func (pq *PriorityQueue[T]) Push(value T) {
	pq.impl.Push(value)
}

// Pop 从 priority queue 弹出元素，并返回
func (pq *PriorityQueue[T]) Pop() T {
	return pq.impl.Pop()
}

type pqImpl[T any] interface {
	Push(value T)
	Pop() T
}

type pqOrdered[T gostl.Ordered] struct {
	PriorityQueue[T]
}

func (q *pqOrdered[T]) Push(value T) {
	heap.PushMinHeap(&q.heap, value)
}

func (q *pqOrdered[T]) Pop() T {
	return heap.PopMinHeap(&q.heap)
}

type pqFunc[T any] struct {
	PriorityQueue[T]
	less gostl.LessFunc[T]
}

func (q *pqFunc[T]) Push(value T) {
	heap.PushMinHeapFunc(&q.heap, value, q.less)
}

func (q *pqFunc[T]) Pop() T {
	return heap.PopMinHeapFunc(&q.heap, q.less)
}
