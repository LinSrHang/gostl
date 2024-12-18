package list

// 双向链表，基于循环链表且带头节点的实现
type DList[T any] struct {
	head   *dListNode[T]
	length int
}

type dListNode[T any] struct {
	prev  *dListNode[T]
	next  *dListNode[T]
	value T
}

// NewDList 构造一个空的双向链表
func NewDList[T any]() DList[T] {
	return DList[T]{}
}

// NewDListInitializer 使用 initializerList 初始化双向链表
func NewDListInitializer[T any](values ...T) DList[T] {
	l := DList[T]{}
	for _, v := range values {
		l.PushBack(v)
	}
	return l
}

// Clear 清空双向链表
func (l *DList[T]) Clear() {
	if l.head != nil {
		l.head.prev = l.head
		l.head.next = l.head
	}
	l.length = 0
}

// Len 返回双向链表的节点数
func (l DList[T]) Len() int {
	return l.length
}

// Empty 返回双向链表是否为空
func (l DList[T]) Empty() bool {
	return l.length == 0
}

// Front 获取双向链表前端元素
func (l DList[T]) Front() T {
	if l.Empty() {
		panic("dList is empty!")
	}
	return l.head.next.value
}

// Back 获取双向链表尾部元素
func (l DList[T]) Back() T {
	if l.Empty() {
		panic("dList is empty!")
	}
	return l.head.prev.value
}

// PushFront 在双向链表头部添加元素
func (l *DList[T]) PushFront(value T) {
	l.ensureHead()
	n := dListNode[T]{prev: l.head, next: l.head.next, value: value}
	l.head.next.prev = &n
	l.head.next = &n
	l.length++
}

// PushBack 在双向链表尾部添加元素
func (l *DList[T]) PushBack(value T) {
	l.ensureHead()
	n := dListNode[T]{prev: l.head.prev, next: l.head, value: value}
	l.head.prev.next = &n
	l.head.prev = &n
	l.length++
}

// PopFront 弹出双向链表头部元素并返回
func (l *DList[T]) PopFront() T {
	if l.Empty() {
		panic("DList.PopFront: empty list")
	}
	node := l.head.next
	value := node.value
	l.head.next = node.next
	l.head.next.prev = l.head
	node.prev = nil
	node.next = nil
	l.length--
	return value
}

// PopBack 弹出双向链表尾部元素并返回
func (l *DList[T]) PopBack() T {
	if l.Empty() {
		panic("DList.PopBack: empty list")
	}
	node := l.head.prev
	value := node.value
	l.head.prev = l.head.prev.prev
	l.head.prev.next = l.head
	node.prev = nil
	node.next = nil
	l.length--
	return value
}

// ForEach 遍历双向链表，并为每个元素执行 f 函数
func (l *DList[T]) ForEach(f func(value *T)) {
	if l.Empty() {
		return
	}
	for node := l.head.next; node != l.head; node = node.next {
		f(&node.value)
	}
}

// ForEachIf 遍历双向链表，并为每个元素执行 f 函数，若其中一个 f 函数返回 false，直接返回
func (l *DList[T]) ForEachIf(f func(value *T) bool) {
	if l.Empty() {
		return
	}
	for node := l.head.next; node != l.head; node = node.next {
		if !f(&node.value) {
			return
		}
	}
}

func (l *DList[T]) ensureHead() {
	if l.head == nil {
		l.head = &dListNode[T]{}
		l.head.prev = l.head
		l.head.next = l.head
	}
}
