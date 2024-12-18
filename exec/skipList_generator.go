package main

import (
	"fmt"
	"os"
)

const pre = `package list

func newSkipListNode[K any, V any](level int, key K, value V) *skipListNode[K, V] {
	switch level {
`

const sec = `	case %d:
		n := struct {
			head  skipListNode[K, V]
			nexts [%d]*skipListNode[K, V]
		}{head: skipListNode[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head
`

const suf = `	default:
		panic("should not reach here")
	}
}`

func main() {
	var skipListMaxLevel int
	fmt.Printf("Enter the maximum level of Skip List: ")
	fmt.Scanf("%d", &skipListMaxLevel)

	file, err := os.Create("./skipList_newNode.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, pre)
	for i := 0; i <= skipListMaxLevel; i++ {
		fmt.Fprintf(file, sec, i, i)
	}
	fmt.Fprint(file, suf)
}
