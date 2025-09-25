package leetcode

/*
在这个 LRUCache 实现中，头尾节点是空的（也称为哑节点或哑指针），主要是为了简化链表的操作逻辑。有以下几个好处：

1、避免了边界条件的判断：在添加或删除节点时，不需要分别处理头节点和尾节点的情况。无论是添加到头部还是删除尾部，都可以统一处理为与相邻的哑节点进行连接。
2、使得头尾指针的更新更容易：当你需要将一个节点移到链表的头部时，只需要更新它的前驱和后继节点，以及头节点的后继节点。同样，删除尾节点也只需要更新尾节点的前驱节点即可。
3、方便维护链表的顺序：由于头尾节点总是存在的，所以你可以保证链表的顺序始终正确，即使在边界情况下（比如链表为空或者只有一个节点）。
4、易于实现双向链表的特性：使用哑节点可以很自然地实现双向链表的特性，例如每个节点都有一个前驱和后继节点，即使是头节点和尾节点。
总的来说，使用哑节点可以让你的代码更加简洁、清晰和易于维护。它是一种常见的技巧，尤其是在处理链表时。

*/

type LRUCache struct {
	cap        int
	m          map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val  int
	pre, next *Node
}

func NewLRUCache(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		cap:  capacity,
		m:    make(map[int]*Node),
		head: head,
		tail: tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.m[key]; ok {
		l.moveToHead(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key, val int) {
	if node, ok := l.m[key]; ok {
		node.val = val
		l.moveToHead(node)
		return
	}
	if len(l.m) == l.cap {
		last := l.removeTail()
		delete(l.m, last.key)
	}

	node := &Node{
		key: key,
		val: val,
	}
	l.m[key] = node
	l.addToHead(node)
	return
}

func (l *LRUCache) addToHead(node *Node) {
	l.head.next.pre = node
	node.next = l.head.next
	node.pre = l.head
	l.head.next = node

}
func (l *LRUCache) removeTail() *Node {
	node := l.tail.pre
	l.removeNode(node)
	return node
}
func (l *LRUCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}
func (l *LRUCache) removeNode(node *Node) {
	pre := node.pre
	next := node.next
	pre.next = next
	next.pre = pre
}
