package vector

import "sync"

type VectorMutex[T any] struct {
	Vector[T]
	mux sync.RWMutex
}

// NewVectorMutex 构造一个空 VectorMutex[T]
func NewVectorMutex[T any]() *VectorMutex[T] {
	return &VectorMutex[T]{
		Vector: NewVector[T](),
		mux:    sync.RWMutex{},
	}
}

// NewVectorMutexWithCapacity 构造一个有初始容量的 VectorMutex[T]
func NewVectorMutexWithCapacity[T any](capacity int) *VectorMutex[T] {
	return &VectorMutex[T]{
		Vector: NewVectorWithCapacity[T](capacity),
		mux:    sync.RWMutex{},
	}
}

// VectorMutexInitializerList 构造一个 VectorMutex[T] 并用 initializerList 初始值
func VectorMutexInitializerList[T any](values ...T) *VectorMutex[T] {
	return &VectorMutex[T]{
		Vector: VectorInitializerList[T](values...),
		mux:    sync.RWMutex{},
	}
}

// Empty 返回 vec 是否为空
func (vec *VectorMutex[T]) Empty() bool {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	return vec.Vector.Empty()
}

// Len 返回 vec 的长度
func (vec *VectorMutex[T]) Len() int {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	return vec.Vector.Len()
}

// Cap 返回 vec 的容量
func (vec *VectorMutex[T]) Cap() int {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	return vec.Vector.Cap()
}

// Clear 清空 vec
func (vec *VectorMutex[T]) Clear() {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.Clear()
}

// Reserve 增加 vec 的容量至 capacity，如果目标容量小于当前容量，不做任何修改
func (vec *VectorMutex[T]) Reserve(capacity int) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.Reserve(capacity)
}

// Shrink 将 vec 的容量调整至当前长度
func (vec *VectorMutex[T]) Shrink() {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.Shrink()
}

// At 获取 vec 特定索引位置的元素
func (vec *VectorMutex[T]) At(idx int) T {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	return vec.Vector.At(idx)
}

// Set 设置 vec 特定索引位置的元素
func (vec *VectorMutex[T]) Set(idx int, value T) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.Set(idx, value)
}

// Append 向 vec 尾部添加若干个元素
func (vec *VectorMutex[T]) Append(values ...T) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.Append(values...)
}

// PopBack 删除 vec 尾部元素并返回该元素
func (vec *VectorMutex[T]) PopBack() T {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	return vec.Vector.PopBack()
}

// Back 返回 vec 尾部元素的拷贝
func (vec *VectorMutex[T]) Back() T {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	return vec.Vector.Back()
}

// Front 返回 vec 头部元素的拷贝
func (vec *VectorMutex[T]) Front() T {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	return vec.Vector.Front()
}

// Insert 在 vec 特定索引位置插入元素
func (vec *VectorMutex[T]) Insert(idx int, values ...T) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.Insert(idx, values...)
}

// RemoveRange 删除 vec 中 [l, r) 之间的元素
func (vec *VectorMutex[T]) RemoveRange(l, r int) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.Vector.RemoveRange(l, r)
}

// Remove 删除 vec 中特定索引位置的元素
func (vec *VectorMutex[T]) Remove(idx int) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.RemoveRange(idx, idx+1)
}

// RemoveLength 删除 vec 中特定索引位置开始的 length 个元素
func (vec *VectorMutex[T]) RemoveLength(idx, length int) {
	vec.mux.Lock()
	defer vec.mux.Unlock()
	vec.RemoveRange(idx, idx+length)
}

// ForEach 遍历 vec，并为每个元素执行 f 函数
func (vec *VectorMutex[T]) ForEach(f func(value *T)) {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	for idx := 0; idx < vec.Len(); idx++ {
		f(&(vec.Vector[idx]))
	}
}

// ForEachIf 遍历 vec，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (vec *VectorMutex[T]) ForEachIf(f func(value *T) bool) {
	vec.mux.RLock()
	defer vec.mux.RUnlock()
	for idx := 0; idx < vec.Len(); idx++ {
		if !f(&(vec.Vector[idx])) {
			return
		}
	}
}
