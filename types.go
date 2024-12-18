package gostl

// 所有底层类型为 int, int8, int16, int32, int64 的类型均能实例化为 Signed
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// 所有底层类型为 uint, uint8, uint16, uint32, uint64 的类型均能实例化为 Unsigned
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// 整数类型
type Integer interface {
	Signed | Unsigned
}

// 所有底层类型为 float32, float64 的类型均能实例化为 Float
type Float interface {
	~float32 | ~float64
}

// 可比较类型（整数、浮点数、字符串、复数）
type Ordered interface {
	Integer | Float | ~string
}

// 数字类型（整数、浮点数、复数）
type Number interface {
	Integer | Float
}

// 返回 a < b 的结果
type LessFunc[T any] func(a, b T) bool

// 返回 c 的哈希
type HashFunc[T any] func(t T) uint64

// 对 a, b 进行比较
//  若 a >  b，返回 1
//  若 a == b，返回 0
//  若 a <  b，返回 -1
type CompareFunc[T any] func(a, b T) int
