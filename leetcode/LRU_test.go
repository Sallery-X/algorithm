package leetcode

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	l := NewLRUCache(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Put(4, 4)
	fmt.Println(l.head.next.next.val)

	x := l.Get(4)
	//fmt.Println(l.values)
	fmt.Println(x)

}
