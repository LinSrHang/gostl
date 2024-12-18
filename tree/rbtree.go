package tree

import "gostl"

const (
	RED   = true
	BLACK = false
)

type rbNode[T any] struct {
	left   *rbNode[T]
	right  *rbNode[T]
	parent *rbNode[T]
	color  bool
	value  T
}

type RBTree[T any] struct {
	root    *rbNode[T]
	nilNode *rbNode[T]
	count   int
	impl    rbImpl[T]
}

// NewRBTree 构造一个可比较类型的红黑树
func NewRBTree[T gostl.Ordered]() *RBTree[T] {
	var zero T
	node := &rbNode[T]{
		left:   nil,
		right:  nil,
		parent: nil,
		color:  BLACK,
		value:  zero,
	}
	rbt := rbTreeOrdered[T]{}
	rbt.root = node
	rbt.nilNode = node
	rbt.count = 0
	rbt.impl = (rbImpl[T])(&rbt)
	return &rbt.RBTree
}

// NewRBTreeFunc 基于比较函数 cmp 构造一个红黑树
func NewRBTreeFunc[T any](less gostl.LessFunc[T]) *RBTree[T] {
	var zero T
	node := &rbNode[T]{
		left:   nil,
		right:  nil,
		parent: nil,
		color:  BLACK,
		value:  zero,
	}
	rbt := rbTreeFunc[T]{}
	rbt.root = node
	rbt.nilNode = node
	rbt.count = 0
	rbt.less = less
	rbt.impl = (rbImpl[T])(&rbt)
	return &rbt.RBTree
}

func (t *RBTree[T]) leftRotate(x *rbNode[T]) {
	if x.right == t.nilNode {
		return
	}

	//          |                                  |
	//          X                                  Y
	//         / \         left rotate            / \
	//        α   Y       ------------->         X   γ
	//           / \                            / \
	//           β  γ                          α  β
	y := x.right
	x.right = y.left
	if y.left != t.nilNode {
		y.left.parent = x
	}
	y.parent = x.parent

	if x.parent == t.nilNode {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *RBTree[T]) rightRotate(x *rbNode[T]) {
	if x.left == t.nilNode {
		return
	}

	//          |                                  |
	//          X                                  Y
	//         / \         right rotate           / \
	//        Y   γ      ------------->         α  X
	//       / \                                    / \
	//      α  β                                   β  γ
	y := x.left
	x.left = y.right
	if y.right != t.nilNode {
		y.right.parent = x
	}
	y.parent = x.parent

	if x.parent == t.nilNode {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}

// Len 返回红黑树的节点数
func (t *RBTree[T]) Len() int {
	return t.count
}

// Insert 往红黑树中插入元素
func (t *RBTree[T]) Insert(value T) {
	t.impl.Insert(&rbNode[T]{
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
		color:  RED,
		value:  value,
	})
}

// InsertOrGet 如果值不存在则插入，值存在则获取并返回
func (t *RBTree[T]) InsertOrGet(value T) T {
	return t.impl.Insert(&rbNode[T]{
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
		color:  RED,
		value:  value,
	}).value
}

// Delete 删除红黑树中的元素并返回
func (t *RBTree[T]) Delete(value T) T {
	return t.delete(&rbNode[T]{
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
		color:  RED,
		value:  value,
	}).value
}

// Search 在红黑树中搜索元素
func (t *RBTree[T]) Search(value T) *rbNode[T] {
	return t.impl.Search(&rbNode[T]{
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
		color:  RED,
		value:  value,
	})
}

// Min 获取整个红黑树的最小值
func (t *RBTree[T]) Min() T {
	var zero T
	x := t.MinSub(t.root)
	if x == t.nilNode {
		return zero
	}
	return x.value
}

// Max 获取整个红黑树的最大值
func (t *RBTree[T]) Max() T {
	var zero T
	x := t.MaxSub(t.root)
	if x == t.nilNode {
		return zero
	}
	return x.value
}

func (t *RBTree[T]) insertFixup(node *rbNode[T]) {
	for node.parent.color { // node.parent.color == RED
		if node.parent == node.parent.parent.left {
			y := node.parent.parent.right
			if y.color { // y.color == RED
				node.parent.color = BLACK
				y.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else { // y.color == BLACK
				if node == node.parent.right {
					node = node.parent
					t.impl.Insert(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.rightRotate(node.parent.parent)
			}
		} else {
			y := node.parent.parent.left
			if y.color { // y.color == RED
				node.parent.color = BLACK
				y.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else { // y.color == BLACK
				if node == node.parent.left {
					node = node.parent
					t.rightRotate(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.leftRotate(node.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

// MinSub 返回红黑树中的以指定节点为根节点的子树的最小值
func (t *RBTree[T]) MinSub(node *rbNode[T]) *rbNode[T] {
	if node == t.nilNode {
		return t.nilNode
	}
	for node.left != t.nilNode {
		node = node.left
	}
	return node
}

// MaxSub 返回红黑树中的以指定节点为根节点的子树的最大值
func (t *RBTree[T]) MaxSub(node *rbNode[T]) *rbNode[T] {
	if node == t.nilNode {
		return t.nilNode
	}
	for node.right != t.nilNode {
		node = node.right
	}
	return node
}

// Get 获取红黑树中的指定节点的值
func (t *RBTree[T]) Get(value T) T {
	ret := t.impl.Search(&rbNode[T]{
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
		color:  RED,
		value:  value,
	})
	if ret == nil {
		var zero T
		return zero
	}
	return ret.value
}

func (t *RBTree[T]) successor(node *rbNode[T]) *rbNode[T] {
	if node == t.nilNode {
		return t.nilNode
	}

	if node.right != t.nilNode {
		return t.MinSub(node.right)
	}

	y := node.parent
	for y != t.nilNode && node == y.right {
		node = y
		y = y.parent
	}
	return y
}

func (t *RBTree[T]) delete(key *rbNode[T]) *rbNode[T] {
	z := t.impl.Search(key)
	if z == t.nilNode {
		return t.nilNode
	}
	ret := &rbNode[T]{
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
		color:  z.color,
		value:  z.value,
	}

	var x, y *rbNode[T]
	if z.left == t.nilNode || z.right == t.nilNode {
		y = z
	} else {
		y = t.successor(z)
	}

	if y.left != t.nilNode {
		x = y.left
	} else {
		x = y.right
	}

	x.parent = y.parent
	if y.parent == t.nilNode {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.value = y.value
	}
	if !y.color { // y.color == BLACK
		t.deleteFixup(x)
	}

	t.count--
	return ret
}

func (t *RBTree[T]) deleteFixup(node *rbNode[T]) {
	for node != t.root && !node.color {
		if node == node.parent.left {
			w := node.parent.right
			if w.color {
				w.color = BLACK
				node.parent.color = RED
				t.leftRotate(node.parent)
				w = node.parent.right
			}
			if !w.left.color && !w.right.color { // w.left.color == BLACK && w.right.color == BLACK
				w.color = RED
				node = node.parent
			} else {
				if !w.right.color {
					w.left.color = BLACK
					w.color = RED
					t.rightRotate(w)
					w = node.parent.right
				}
				w.color = node.parent.color
				node.parent.color = BLACK
				w.right.color = BLACK
				t.leftRotate(node.parent)
				node = t.root
			}
		} else {
			w := node.parent.left
			if w.color {
				w.color = BLACK
				node.parent.color = RED
				t.rightRotate(node.parent)
				w = node.parent.left
			}
			if !w.left.color && !w.right.color { // w.left.color == BLACK && w.right.color == BLACK
				w.color = RED
				node = node.parent
			} else {
				if !w.left.color {
					w.right.color = BLACK
					w.color = RED
					t.leftRotate(w)
					w = node.parent.left
				}
				w.color = node.parent.color
				node.parent.color = BLACK
				w.left.color = BLACK
				t.rightRotate(node.parent)
				node = t.root
			}
		}
	}
	node.color = BLACK
}

type rbImpl[T any] interface {
	Insert(node *rbNode[T]) *rbNode[T]
	Search(node *rbNode[T]) *rbNode[T]
}

type rbTreeOrdered[T gostl.Ordered] struct {
	RBTree[T]
}

func (t *rbTreeOrdered[T]) Insert(node *rbNode[T]) *rbNode[T] {
	x := t.root
	y := t.nilNode

	for x != t.nilNode {
		y = x
		if node.value < x.value {
			x = x.left
		} else if x.value < node.value {
			x = x.right
		} else {
			return x
		}
	}

	node.parent = y
	if y == t.nilNode {
		t.root = node
	} else if node.value < y.value {
		y.left = node
	} else {
		y.right = node
	}

	t.count++
	t.insertFixup(node)
	return node
}

func (t *rbTreeOrdered[T]) Search(node *rbNode[T]) *rbNode[T] {
	p := t.root

	for p != t.nilNode {
		if p.value < node.value {
			p = p.right
		} else if node.value < p.value {
			p = p.left
		} else {
			break
		}
	}

	return p
}

type rbTreeFunc[T any] struct {
	RBTree[T]
	less gostl.LessFunc[T]
}

func (t *rbTreeFunc[T]) Insert(node *rbNode[T]) *rbNode[T] {
	x := t.root
	y := t.nilNode

	for x != t.nilNode {
		y = x
		if t.less(node.value, x.value) {
			x = x.left
		} else if t.less(x.value, node.value) {
			x = x.right
		} else {
			return x
		}
	}

	node.parent = y
	if y == t.nilNode {
		t.root = node
	} else if t.less(node.value, y.value) {
		y.left = node
	} else {
		y.right = node
	}

	t.count++
	t.insertFixup(node)
	return node
}

func (t *rbTreeFunc[T]) Search(node *rbNode[T]) *rbNode[T] {
	p := t.root

	for p != t.nilNode {
		if t.less(p.value, node.value) {
			p = p.right
		} else if t.less(node.value, p.value) {
			p = p.left
		} else {
			break
		}
	}

	return p
}
