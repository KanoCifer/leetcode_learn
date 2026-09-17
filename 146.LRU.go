package leetcodelearn

type LRUCache struct {
	Cache    map[int]*Node
	Capacity int
	Head     *Node
	Tail     *Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, head, nil}
	head.Next = tail
	return LRUCache{
		Cache:    make(map[int]*Node),
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	cur := this.Cache[key]
	if cur != nil {
		this.moveToHead(cur)
		return cur.Val
	}

	return -1
}

func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) addToHead(node *Node) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.Cache[key]; ok {
		// 命中:更新值并升到 head
		node.Val = value
		this.removeNode(node)
		this.addToHead(node)
		return
	}

	if len(this.Cache) >= this.Capacity {
		victim := this.Tail.Prev
		this.removeNode(victim)
		delete(this.Cache, victim.Key)
	}
	node := &Node{Key: key, Val: value}
	this.Cache[key] = node
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
