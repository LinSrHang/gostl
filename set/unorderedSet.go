package set

// 无序集合，基于 golang map 实现
type UnorderedSet[K comparable] map[K]struct{}

// NewUnorderedSet 构造一个空的无序集合
func NewUnorderedSet[K comparable]() UnorderedSet[K] {
	return make(UnorderedSet[K])
}

// NewUnorderedSetInitializerList 使用 initializerList 初始化无序集合
func NewUnorderedSetInitializerList[K comparable](keys ...K) UnorderedSet[K] {
	s := NewUnorderedSet[K]()
	s.QuickInsert(keys...)
	return s
}

// Empty 判断无序集合是否为空
func (s UnorderedSet[K]) Empty() bool {
	return len(s) == 0
}

// Len 获取无序集合中元素的数量
func (s UnorderedSet[K]) Len() int {
	return len(s)
}

// Clear 清空无序集合
func (s *UnorderedSet[K]) Clear() {
	for k := range *s {
		delete(*s, k)
	}
}

// Exist 判断无序集合中是否存在指定键
func (s UnorderedSet[K]) Exist(key K) bool {
	_, ok := s[key]
	return ok
}

// Insert 向无序集合中插入键，返回是否成功插入
//  忽略返回值时建议使用 QuickInsert
func (s *UnorderedSet[K]) Insert(key K) bool {
	oldLen := len(*s)
	(*s)[key] = struct{}{}
	return len(*s) > oldLen
}

// InsertN 向无序集合中批量插入键，返回插入成功的键的数量
//  忽略返回值时建议使用 QuickInsert
func (s *UnorderedSet[K]) InsertN(keys ...K) int {
	oldLen := len(*s)
	for _, key := range keys {
		(*s)[key] = struct{}{}
	}
	return len(*s) - oldLen
}

// QuickInsert 快速插入若干键，跳过检查是否插入成功的实现
//  使用 Insert 和 InsertN 且忽略返回值时建议使用 QuickInsert
func (s *UnorderedSet[K]) QuickInsert(keys ...K) {
	for _, key := range keys {
		(*s)[key] = struct{}{}
	}
}

// Remove 从无序集合中移除指定键，返回移除前是否存在
//  忽略返回值时建议使用 QuickRemove
func (s *UnorderedSet[K]) Remove(key K) bool {
	_, ok := (*s)[key]
	delete(*s, key)
	return ok
}

// RemoveN 从无序数组中移除若干键，返回成功移除的数量
//  忽略返回值时建议使用 QuickRemove
func (s *UnorderedSet[K]) RemoveN(keys ...K) int {
	oldLen := len(*s)
	for _, key := range keys {
		delete(*s, key)
	}
	return oldLen - len(*s)
}

// QuickRemove 快速移除若干键，跳过检查是否移除成功的实现
//  使用 Remove 和 RemoveN 且忽略返回值时建议使用 QuickRemove
func (s *UnorderedSet[K]) QuickRemove(keys ...K) {
	for _, key := range keys {
		delete(*s, key)
	}
}

// Keys 返回无序集合中所有键
func (s UnorderedSet[K]) Keys() []K {
	keys := make([]K, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}

// ForEach 遍历无序集合，并为每个元素执行 f 函数
func (s *UnorderedSet[K]) ForEach(f func(key *K)) {
	for key := range *s {
		f(&key)
	}
}

// ForEachIf 遍历无序集合，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (s *UnorderedSet[K]) ForEachIf(f func(key *K) bool) {
	for key := range *s {
		if !f(&key) {
			return
		}
	}
}

// InsertSet 将 other 集合插入到当前集合
func (s *UnorderedSet[K]) InsertSet(other *UnorderedSet[K]) {
	for key := range *other {
		(*s)[key] = struct{}{}
	}
}

// Union 返回当前集合和 other 集合的合并结果，不修改当前集合
func (s UnorderedSet[K]) Union(other *UnorderedSet[K]) UnorderedSet[K] {
	res := NewUnorderedSet[K]()
	res.InsertSet(&s)
	res.InsertSet(other)
	return res
}

// orderSet 比较两个无序集合的长度
func orderSet[K comparable](a, b *UnorderedSet[K]) (small, large *UnorderedSet[K]) {
	if a.Len() < b.Len() {
		return a, b
	}
	return b, a
}

// Intersection 返回当前集合和 other 集合的交集，不修改当前集合
func (s UnorderedSet[K]) Intersection(other *UnorderedSet[K]) UnorderedSet[K] {
	res := NewUnorderedSet[K]()
	small, large := orderSet(&s, other)
	for key := range *small {
		if large.Exist(key) {
			res.QuickInsert(key)
		}
	}
	return res
}

// Difference 返回当前集合和 other 集合的差集，不修改当前集合
func (s UnorderedSet[K]) Difference(other *UnorderedSet[K]) UnorderedSet[K] {
	res := NewUnorderedSet[K]()
	for key := range s {
		if !other.Exist(key) {
			res.QuickInsert(key)
		}
	}
	return res
}

// Disjoint 判断当前集合与 other 集合是否不相交
func (s UnorderedSet[K]) Disjoint(other UnorderedSet[K]) bool {
	small, large := orderSet(&s, &other)
	for key := range *small {
		if large.Exist(key) {
			return false
		}
	}
	return true
}

// Subset 判断当前集合是否为 other 集合的子集
func (s UnorderedSet[K]) Subset(other UnorderedSet[K]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for key := range s {
		if !other.Exist(key) {
			return false
		}
	}
	return true
}

// Superset 判断当前集合是否为 other 集合的超集
func (s UnorderedSet[K]) Superset(other UnorderedSet[K]) bool {
	return other.Subset(s)
}
