package list

import (
	"gostl"
	"math/bits"
	"math/rand"
	"time"
)

const skipListMaxLevel = 40

type SkipList[K any, V any] struct {
	level      int                // 当前层级
	length     int                // 跳表中拥有的元素总数
	head       skipListNode[K, V] // head.next[level] 为每一层级的头节点
	prevsCache []*skipListNode[K, V]
	rander     *rand.Rand
	impl       skipListImpl[K, V]
}

// NewSkipList 构造一个空的跳表
func NewSkipList[K gostl.Ordered, V any]() *SkipList[K, V] {
	l := skipListOrdered[K, V]{}
	l.init()
	l.impl = (skipListImpl[K, V])(&l)
	return &l.SkipList
}

// NewSkipListFromMap 构造一个跳表，并从 map 中插入元素
func NewSkipListFromMap[K gostl.Ordered, V any](m map[K]V) *SkipList[K, V] {
	l := NewSkipList[K, V]()
	for k, v := range m {
		l.Insert(k, v)
	}
	return l
}

// NewSkipListFunc 构造一个跳表，并使用提供的 keyComp 作为 key 排序函数
func NewSkipListFunc[K any, V any](keyCmp gostl.CompareFunc[K]) *SkipList[K, V] {
	l := skipListFunc[K, V]{}
	l.init()
	l.keyCmp = keyCmp
	l.impl = (skipListImpl[K, V])(&l)
	return &l.SkipList
}

// Empty 判断跳表是否为空
func (l *SkipList[K, V]) Empty() bool {
	return l.length == 0
}

// Len 获取跳表中元素的数量
func (l *SkipList[K, V]) Len() int {
	return l.length
}

// Clear 清空跳表
func (l *SkipList[K, V]) Clear() {
	for i := range l.head.next {
		l.head.next[i] = nil
	}
	l.level = 1
	l.length = 0
}

// Insert 往跳表中插入一对键值对
//
//	如果键已经存在，则更新对应的值
func (l *SkipList[K, V]) Insert(key K, value V) {
	node, prevs := l.impl.findInsertPoint(key)
	if node != nil {
		// 键已存在，仅更新值
		node.value = value
		return
	}

	level := l.randomLevel()
	node = newSkipListNode(level, key, value)

	for i := 0; i < min(level, l.level); i++ {
		node.next[i] = prevs[i].next[i]
		prevs[i].next[i] = node
	}

	if level > l.level {
		for i := l.level; i < level; i++ {
			l.head.next[i] = node
		}
		l.level = level
	}
	l.length++
}

// Find 返回指定键对应的值的引用，如果键不存在，返回 nil
func (l *SkipList[K, V]) Find(key K) *V {
	node := l.impl.findNode(key)
	if node == nil {
		return nil
	}
	return &node.value
}

// Exist 判断跳表中是否存在指定键
func (l *SkipList[K, V]) Exist(key K) bool {
	return l.impl.findNode(key) != nil
}

// Remove 删除跳表中指定键值对，如果键值对不存在，返回 false
func (l *SkipList[K, V]) Remove(key K) bool {
	node, prevs := l.impl.findRemovePoint(key)
	if node == nil {
		return false
	}
	for i, v := range node.next {
		prevs[i].next[i] = v
	}
	for l.level > 1 && l.head.next[l.level-1] == nil {
		l.level--
	}
	l.length--
	return true
}

// ForEach 遍历跳表，并为每个元素执行 f 函数
func (l *SkipList[K, V]) ForEach(f func(key K, value *V)) {
	for e := l.head.next[0]; e != nil; e = e.next[0] {
		f(e.key, &e.value)
	}
}

// ForEachIf 遍历跳表，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (l *SkipList[K, V]) ForEachIf(f func(key K, value *V) bool) {
	for e := l.head.next[0]; e != nil; e = e.next[0] {
		if !f(e.key, &e.value) {
			return
		}
	}
}

type skipListNode[K any, V any] struct {
	key   K
	value V
	next  []*skipListNode[K, V] // 后驱指针
}

type skipListImpl[K any, V any] interface {
	findNode(key K) *skipListNode[K, V]
	lowerBound(key K) *skipListNode[K, V]
	upperBound(key K) *skipListNode[K, V]
	findInsertPoint(key K) (*skipListNode[K, V], []*skipListNode[K, V])
	findRemovePoint(key K) (*skipListNode[K, V], []*skipListNode[K, V])
}

func (l *SkipList[K, V]) init() {
	l.level = 1
	l.rander = rand.New(rand.NewSource(time.Now().Unix()))
	l.prevsCache = make([]*skipListNode[K, V], skipListMaxLevel)
	l.head.next = make([]*skipListNode[K, V], skipListMaxLevel)
}

func (l *SkipList[K, V]) randomLevel() int {
	total := uint64(1)<<uint64(skipListMaxLevel) - 1 // 2^n -1
	k := l.rander.Uint64() & total
	level := skipListMaxLevel - bits.Len64(k) + 1
	for level > 3 && 1<<(level-3) > l.length {
		level--
	}
	return level
}

type skipListOrdered[K gostl.Ordered, V any] struct {
	SkipList[K, V]
}

func (l *skipListOrdered[K, V]) findNode(key K) *skipListNode[K, V] {
	return l.doFindOne(key, true)
}

func (l *skipListOrdered[K, V]) doFindOne(key K, eq bool) *skipListNode[K, V] {
	prev := &l.head
	for i := l.level - 1; i >= 0; i-- {
		for cur := prev.next[i]; cur != nil; cur = cur.next[i] {
			if cur.key == key {
				return cur
			}
			if cur.key > key {
				break
			}
			prev = cur
		}
	}
	if eq {
		return nil
	}
	return prev.next[0]
}

func (l *skipListOrdered[K, V]) lowerBound(key K) *skipListNode[K, V] {
	return l.doFindOne(key, false)
}

func (l *skipListOrdered[K, V]) upperBound(key K) *skipListNode[K, V] {
	node := l.lowerBound(key)
	if node != nil && node.key == key {
		return node.next[0]
	}
	return node
}

func (l *skipListOrdered[K, V]) findInsertPoint(key K) (*skipListNode[K, V], []*skipListNode[K, V]) {
	prevs := l.prevsCache[0:l.level]
	prev := &l.head
	for i := l.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if next.key == key {
				return next, nil
			}
			if next.key > key {
				break
			}
			prev = next
		}
		prevs[i] = prev
	}
	return nil, prevs
}

func (l *skipListOrdered[K, V]) findRemovePoint(key K) (*skipListNode[K, V], []*skipListNode[K, V]) {
	prevs := l.findPrevNodes(key)
	node := prevs[0].next[0]
	if node == nil || node.key != key {
		return nil, nil
	}
	return node, prevs
}

func (l *skipListOrdered[K, V]) findPrevNodes(key K) []*skipListNode[K, V] {
	prevs := l.prevsCache[0:l.level]
	prev := &l.head
	for i := l.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if next.key >= key {
				break
			}
			prev = next
		}
		prevs[i] = prev
	}
	return prevs
}

type skipListFunc[K any, V any] struct {
	SkipList[K, V]
	keyCmp gostl.CompareFunc[K]
}

func (l *skipListFunc[K, V]) findNode(key K) *skipListNode[K, V] {
	node := l.lowerBound(key)
	if node != nil && l.keyCmp(node.key, key) == 0 {
		return node
	}
	return nil
}

func (l *skipListFunc[K, V]) lowerBound(key K) *skipListNode[K, V] {
	var prev = &l.head
	for i := l.level - 1; i >= 0; i-- {
		cur := prev.next[i]
		for ; cur != nil; cur = cur.next[i] {
			cmpRet := l.keyCmp(cur.key, key)
			if cmpRet == 0 {
				return cur
			}
			if cmpRet > 0 {
				break
			}
			prev = cur
		}
	}
	return prev.next[0]
}

func (l *skipListFunc[K, V]) upperBound(key K) *skipListNode[K, V] {
	node := l.lowerBound(key)
	if node != nil && l.keyCmp(node.key, key) == 0 {
		return node.next[0]
	}
	return node
}

func (l *skipListFunc[K, V]) findInsertPoint(key K) (*skipListNode[K, V], []*skipListNode[K, V]) {
	prevs := l.prevsCache[0:l.level]
	prev := &l.head
	for i := l.level - 1; i >= 0; i-- {
		for cur := prev.next[i]; cur != nil; cur = cur.next[i] {
			r := l.keyCmp(cur.key, key)
			if r == 0 {
				// The key is already existed, prevs are useless because no new node insertion.
				// stop searching.
				return cur, nil
			}
			if r > 0 {
				break
			}
			prev = cur
		}
		prevs[i] = prev
	}
	return nil, prevs
}

func (l *skipListFunc[K, V]) findRemovePoint(key K) (*skipListNode[K, V], []*skipListNode[K, V]) {
	prevs := l.findPrevNodes(key)
	node := prevs[0].next[0]
	if node == nil || l.keyCmp(node.key, key) != 0 {
		return nil, nil
	}
	return node, prevs
}

func (l *skipListFunc[K, V]) findPrevNodes(key K) []*skipListNode[K, V] {
	prevs := l.prevsCache[0:l.level]
	prev := &l.head
	for i := l.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if l.keyCmp(next.key, key) >= 0 {
				break
			}
			prev = next
		}
		prevs[i] = prev
	}
	return prevs
}
