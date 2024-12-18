package stack

import "gostl/vector"

// Stack 栈
type Stack[T any] struct {
	elements vector.Vector[T]
}

// NewStack 构造一个空 Stack[T]
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: vector.NewVector[T]()}
}

// NewStackWithCapacity 构造一个指定容量的 Stack[T]
func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{elements: vector.NewVectorWithCapacity[T](capacity)}
}

// StackInitializerList 构造一个 Stack[T] 并用 initializerList 依次压栈，栈容量等于 initializerList 的长度
func StackInitializerList[T any](values ...T) *Stack[T] {
	return &Stack[T]{elements: vector.VectorInitializerList[T](values...)}
}

// Empty 返回 stk 是否为空（长度）
func (stk *Stack[T]) Empty() bool {
	return stk.elements.Empty()
}

// Len 返回 stk 的长度
func (stk *Stack[T]) Len() int {
	return stk.elements.Len()
}

// Cap 返回 stk 的容量
func (stk *Stack[T]) Cap() int {
	return stk.elements.Cap()
}

// Clear 清空 stk
func (stk *Stack[T]) Clear() {
	stk.elements.Clear()
}

// Reserve 增加 stk 的容量至 capacity，如果目标容量小于当前容量，不做任何修改
func (stk *Stack[T]) Reserve(capacity int) {
	stk.elements.Reserve(capacity)
}

// Shrink 将 stk 的容量调整至当前长度
func (stk *Stack[T]) Shrink() {
	stk.elements.Shrink()
}

// Top 返回 stk 的栈顶元素
func (stk *Stack[T]) Top() T {
	return stk.elements.Back()
}

// Push 依次压入若干个元素 values 到 stk 的栈顶
func (stk *Stack[T]) Push(values ...T) {
	stk.elements.Append(values...)
}

// Pop 出 stk 的栈顶元素并返回
func (stk *Stack[T]) Pop() T {
	return stk.elements.PopBack()
}

// ForEach 遍历 stk，并为每个元素执行 f 函数
func (stk *Stack[T]) ForEach(f func(value *T)) {
	stk.elements.ForEach(f)
}

// ForEachIf 遍历 stk，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (stk *Stack[T]) ForEachIf(f func(value *T) bool) {
	stk.elements.ForEachIf(f)
}
