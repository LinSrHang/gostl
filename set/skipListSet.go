package set

import (
	"gostl"
	"gostl/list"
)

// SkipListSet 跳表实现的有序集合
type SkipListSet[K any] list.SkipList[K, struct{}]

// NewSkipListSet 构造一个空的有序集合
func NewSkipListSet[K gostl.Ordered]() *SkipListSet[K] {
	return (*SkipListSet[K])(list.NewSkipList[K, struct{}]())
}

// NewSkipListSetFunc 构造一个空的有序集合，并指定一个自定义的键值比较函数
func NewSkipListSetFunc[K any](cmp gostl.CompareFunc[K]) *SkipListSet[K] {
	return (*SkipListSet[K])(list.NewSkipListFunc[K, struct{}](cmp))
}

// NewSkipListSetInitializer 构造一个空的有序集合并使用 initializerList 初始化
func NewSkipListInitializer[K gostl.Ordered](elements ...K) *SkipListSet[K] {
	s := NewSkipListSet[K]()
	for i := range elements {
		s.Insert(elements[i])
	}
	return s
}

func (s *SkipListSet[K]) asMap() *list.SkipList[K, struct{}] {
	return (*list.SkipList[K, struct{}])(s)
}

// Empty 判断有序集合是否为空
func (s *SkipListSet[K]) Empty() bool {
	return s.asMap().Empty()
}

// Len 获取有序集合中元素的数量
func (s *SkipListSet[K]) Len() int {
	return s.asMap().Len()
}

// Clear 清空有序集合
func (s *SkipListSet[K]) Clear() {
	s.asMap().Clear()
}

// Exist 判断有序集合中是否存在指定键
func (s *SkipListSet[K]) Exist(key K) bool {
	return s.asMap().Exist(key)
}

// Insert 向有序集合中插入指定键，如果键已存在则返回 false
func (s *SkipListSet[K]) Insert(key K) bool {
	oldLen := s.Len()
	s.asMap().Insert(key, struct{}{})
	return s.Len() > oldLen
}

// InsertN 向有序集合中插入若干键，返回成功插入的数量
func (s *SkipListSet[K]) InsertN(keys ...K) int {
	oldLen := s.Len()
	for i := range keys {
		s.Insert(keys[i])
	}
	return s.Len() - oldLen
}

// Remove 删除有序集合中指定键，返回是否删除成功
func (s *SkipListSet[K]) Remove(keys K) bool {
	return s.asMap().Remove(keys)
}

// RemoveN 删除有序集合中若干键，返回成功删除的数量
func (s *SkipListSet[K]) RemoveN(keys ...K) int {
	oldLen := s.Len()
	for i := range keys {
		s.Remove(keys[i])
	}
	return oldLen - s.Len()
}

// Keys 获取有序集合中所有键的切片，并按升序排序
func (s *SkipListSet[K]) Keys() []K {
	keys := make([]K, 0, s.Len())
	s.ForEach(func(k K) {
		keys = append(keys, k)
	})
	return keys
}

// ForEach 遍历有序集合，并为每个元素执行 f 函数
func (s *SkipListSet[K]) ForEach(f func(key K)) {
	s.asMap().ForEach(func(k K, s *struct{}) {
		f(k)
	})
}

// ForEachIf 遍历有序集合，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (s *SkipListSet[K]) ForEachIf(f func(key K) bool) {
	s.asMap().ForEachIf(func(k K, s *struct{}) bool {
		return f(k)
	})
}
