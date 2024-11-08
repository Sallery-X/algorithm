package leetcode

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	l := NewLruCache(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Put(4, 4)
	fmt.Println(l.head.Value)

	x := l.Get(1)
	//fmt.Println(l.values)
	fmt.Println(x)

}
