package vector

import "gostl"

// 为 []T 起别名为 Vector[T]
type Vector[T any] []T

// NewVector 构造一个空 Vector[T]
func NewVector[T any]() Vector[T] {
	return (Vector[T])([]T{})
}

// NewVectorWithCapacity 构造一个有初始容量的 Vector[T]
func NewVectorWithCapacity[T any](capacity int) Vector[T] {
	return (Vector[T])(make([]T, 0, capacity))
}

// VectorInitializerList 构造一个 Vector[T] 并用 initializerList 初始值
func VectorInitializerList[T any](values ...T) Vector[T] {
	return (Vector[T])(values)
}

// Empty 返回 vec 是否为空（长度）
func (vec *Vector[T]) Empty() bool {
	return len(*vec) == 0
}

// Len 返回 vec 的长度
func (vec *Vector[T]) Len() int {
	return len(*vec)
}

// Cap 返回 vec 的容量
func (vec *Vector[T]) Cap() int {
	return cap(*vec)
}

// Clear 清空
func (vec *Vector[T]) Clear() {
	gostl.FillZero(*vec)
	*vec = (*vec)[:0]
}

// Reserve 增加 vec 的容量至 capacity，如果目标容量小于当前容量，不做任何修改
func (vec *Vector[T]) Reserve(capacity int) {
	if cap(*vec) < capacity {
		t := make([]T, len(*vec), capacity)
		copy(t, *vec)
		*vec = t
	}
}

// Shrink 将 vec 的容量调整至当前长度
func (vec *Vector[T]) Shrink() {
	if cap(*vec) > len(*vec) {
		t := make([]T, len(*vec))
		copy(t, *vec)
		*vec = t
	}
}

// At 返回索引对应的值
func (vec *Vector[T]) At(idx int) T {
	return (*vec)[idx]
}

// Set 设置索引位置对应的值
func (vec *Vector[T]) Set(idx int, value T) {
	(*vec)[idx] = value
}

// Append 向 vec 尾部添加若干个元素
func (vec *Vector[T]) Append(values ...T) {
	*vec = append(*vec, values...)
}

// PopBack 删除 vec 尾部元素并返回该元素
func (vec *Vector[T]) PopBack() T {
	var zero T
	back := (*vec)[vec.Len()-1]
	(*vec)[vec.Len()-1] = zero
	*vec = (*vec)[:vec.Len()-1]
	return back
}

// Back 返回 vec 尾部元素的拷贝
func (vec *Vector[T]) Back() T {
	return (*vec)[vec.Len()-1]
}

// Front 返回 vec 头部元素的拷贝
func (vec *Vector[T]) Front() T {
	return (*vec)[0]
}

// Insert 向 vec 特定索引位置插入若干个元素
func (vec *Vector[T]) Insert(idx int, values ...T) {
	vec1 := *vec
	total := vec1.Len() + len(values)
	if total <= vec1.Cap() {
		vec2 := vec1[:total]
		copy(vec2[idx+len(values):], vec1[idx:])
		copy(vec2[idx:], values)
		*vec = vec2
		return
	}
	vec2 := make([]T, total)
	copy(vec2, vec1[:idx])
	copy(vec2[idx:], values)
	copy(vec2[idx+len(values):], vec1[idx:])
	*vec = vec2
}

// RemoveRange 删除 vec 中 [l, r) 之间的元素
func (vec *Vector[T]) RemoveRange(l, r int) {
	oldVec := *vec
	*vec = append((*vec)[:l], (*vec)[r:]...)
	gostl.FillZero(oldVec[vec.Len():])
}

// Remove 删除 vec 中特定索引位置的元素
func (vec *Vector[T]) Remove(idx int) {
	vec.RemoveRange(idx, idx+1)
}

// RemoveLength 删除 vec 中特定索引位置开始的 length 个元素
func (vec *Vector[T]) RemoveLength(idx, length int) {
	vec.RemoveRange(idx, idx+length)
}

// ForEach 遍历 vec，并为每个元素执行 f 函数
func (vec *Vector[T]) ForEach(f func(value *T)) {
	for idx := 0; idx < vec.Len(); idx++ {
		f(&(*vec)[idx])
	}
}

// ForEach 遍历 vec，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (vec *Vector[T]) ForEachIf(f func(value *T) bool) {
	for idx := 0; idx < vec.Len(); idx++ {
		if !f(&(*vec)[idx]) {
			return
		}
	}
}

// Swap 交换 vec 中两个元素的值
func (vec *Vector[T]) Swap(i, j int) {
	(*vec)[i], (*vec)[j] = (*vec)[j], (*vec)[i]
}

// Reverse 反转 vec 元素的顺序
func (vec *Vector[T]) Reverse() {
	length := vec.Len()
	for i := 0; i < (length >> 1); i++ {
		vec.Swap(i, length-1-i)
	}
}
