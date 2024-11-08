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
	capacity   int
	m          map[int]*Node
	head, tail *Node
}

type Node struct {
	Key       int
	Value     int
	Pre, Next *Node
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.m[key]; ok {
		l.moveToHead(v)
		return v.Value
	}
	return -1
}

func (l *LRUCache) moveToHead(node *Node) {
	l.deleteNode(node)
	l.addToHead(node)
}

func (l *LRUCache) deleteNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LRUCache) removeTail() int {
	node := l.tail.Pre
	l.deleteNode(node)
	return node.Key
}

func (l *LRUCache) addToHead(node *Node) {
	l.head.Next.Pre = node
	node.Next = l.head.Next
	node.Pre = l.head
	l.head.Next = node
}

func (l *LRUCache) Put(key int, value int) {
	if v, ok := l.m[key]; ok {
		v.Value = value
		l.moveToHead(v)
		return
	}

	if l.capacity == len(l.m) {
		rmKey := l.removeTail()
		delete(l.m, rmKey)
	}

	newNode := &Node{Key: key, Value: value}
	l.addToHead(newNode)
	l.m[key] = newNode
}

// 头尾都是空的
func NewLruCache(capacity int) *LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return &LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}
