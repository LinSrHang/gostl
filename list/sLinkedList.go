package list

// SList 单向链表
type SList[T any] struct {
	head   *sListNode[T] // 头指针
	tail   *sListNode[T] // 尾指针
	length int           // 总节点数
}

type sListNode[T any] struct {
	next  *sListNode[T] // 后驱指针
	value T             // 当前节点元素值
}

func NewSList[T any]() SList[T] {
	return SList[T]{}
}

// LinkedListInitializerList 使用 initializerList 初始化单向链表
func SListInitializerList[T any](values ...T) SList[T] {
	l := SList[T]{}
	l.PushBack(values...)
	return l
}

// Empty 返回单向链表是否为空
func (l SList[T]) Empty() bool {
	return l.length == 0
}

// Len 返回单向链表所拥有的节点数
func (l SList[T]) Len() int {
	return l.length
}

func (l *SList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Front 获取单向链表头部元素，若链表为空则 panic
func (l SList[T]) Front() T {
	if l.Empty() {
		panic("sList is empty!")
	}
	return l.head.value
}

// Back 获取单向链表尾部元素，若链表为空则 panic
func (l SList[T]) Back() T {
	if l.Empty() {
		panic("sList is empty!")
	}
	return l.tail.value
}

// PushFront 向单向链表头部添加 values 列表
// 	eg:
//		before: l == 0 -> 1 -> 2
//		do:     l.PushFront(3, 4, 5)
//		after:  l == 3 -> 4 -> 5 -> 0 -> 1 -> 2
func (l *SList[T]) PushFront(values ...T) {
	for idx := len(values) - 1; idx >= 0; idx-- {
		node := sListNode[T]{next: l.head, value: values[idx]}
		l.head = &node
		if l.head == nil {
			l.head = &node
		}
	}
	l.length += len(values)
}

// PushBack 向单向链表尾部依次添加若干元素
func (l *SList[T]) PushBack(values ...T) {
	for idx := range values {
		node := sListNode[T]{next: nil, value: values[idx]}
		if l.tail != nil {
			l.tail.next = &node
		}
		l.tail = &node
		if l.head == nil {
			l.head = &node
		}
	}
	l.length += len(values)
}

// PopFront 移除并返回单向链表头部元素，若链表为空则 panic
func (l *SList[T]) PopFront() T {
	if l.Empty() {
		panic("sList is empty!")
	}

	value := l.head.value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.length--
	return value
}

// PopBack 移除并返回单向链表尾部元素，若链表为空则 panic
func (l *SList[T]) PopBack() T {
	if l.Empty() {
		panic("sList is empty!")
	}

	if l.length == 1 {
		value := l.head.value
		l.head = nil
		l.tail = nil
		l.length = 0
		return value
	}
	value := l.tail.value
	ptr := l.head
	for ptr.next != l.tail {
		ptr = ptr.next
	}
	ptr.next = nil
	l.tail = ptr
	l.length--
	return value
}

// Reverse 反转单向链表
func (l *SList[T]) Reverse() {
	var head, tail *sListNode[T]
	for node := l.head; node != nil; {
		next := node.next
		node.next = head
		head = node
		if tail == nil {
			tail = node
		}
		node = next
	}
	l.head = head
	l.tail = tail
}

// ToVector 将单向链表转换为数组
func (l SList[T]) ToVector() []T {
	vec := (make([]T, 0, l.Len()))
	for node := l.head; node != nil; node = node.next {
		vec = append(vec, node.value)
	}
	return vec
}

// ForEach 遍历单向链表，并为每个元素执行 f 函数
func (l *SList[T]) ForEach(f func(value *T)) {
	for node := l.head; node != nil; node = node.next {
		f(&node.value)
	}
}

// ForEachIf 遍历单向链表，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (l *SList[T]) ForEachIf(f func(value *T) bool) {
	for node := l.head; node != nil; node = node.next {
		if !f(&node.value) {
			return
		}
	}
}
