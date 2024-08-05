package bytedance

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
