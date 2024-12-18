package vector

import (
	"fmt"
	"testing"
)

func Test_NewVector(t *testing.T) {
	fmt.Println(NewVector[int]())
	fmt.Println(make(Vector[int], 1))
	fmt.Println(make(Vector[int], 1, 2))
	v := Vector[int]{1, 2, 3}
	fmt.Println(v)
}

func Test_NewVectorWithCapacity(t *testing.T) {
	v := NewVectorWithCapacity[int](10)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
}

func Test_VectorInitializerList(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	v = VectorInitializerList[int]([]int{4, 5, 6}...)
	fmt.Println(v)
}

func Test_Cap(t *testing.T) {
	v := NewVectorWithCapacity[int](10)
	v.Append(1)
	fmt.Println(v.Len())
	fmt.Println(v.Empty())
	fmt.Println(v.Cap())
}

func Test_Clear(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	fmt.Println(v.Empty())
	v.Clear()
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	fmt.Println(v.Empty())
}

func Test_Reserve(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	v.Reserve(5)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
}

func Test_Shrink(t *testing.T) {
	v := NewVectorWithCapacity[int](10)
	v.Append(1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	v.Shrink()
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
}

func Test_At_Set(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.At(0))
	v[0] = 2
	fmt.Println(v)
	fmt.Println(v.At(0))
	v.Set(3, 0)
}

func Test_Append(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	v.Append(4, 5, 6)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
}

func Test_PopBack(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	v.PopBack()
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
}

func Test_Back(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()
	v := VectorInitializerList[int](1)
	fmt.Println(v.Back())
	v.PopBack()
	fmt.Println(v.Back())
}

func Test_Insert(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
	v.Insert(2, 0, 1, 2)
	fmt.Println(v)
	fmt.Println(v.Len())
	fmt.Println(v.Cap())
}

func Test_Remove(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3, 4, 5)
	oldV := v
	v.Remove(1)
	fmt.Println(v)
	fmt.Println(oldV)
	fmt.Println(oldV[1])
}

func Test_RemoveRange(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3, 4, 5)
	oldV := v
	v.RemoveRange(1, 3)
	fmt.Println(v)
	fmt.Println(oldV)
	fmt.Println(oldV.Len())
}

func Test_RemoveLength(t *testing.T) {
	v := VectorInitializerList[int](1, 2, 3, 4, 5)
	oldV := v
	v.RemoveLength(1, 2)
	fmt.Println(v)
	fmt.Println(oldV)
	fmt.Println(oldV.Len())
}
