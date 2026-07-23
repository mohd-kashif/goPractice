package agodaprep

type Node struct {
	Key, Value int
	Prev, Next *Node
}

type LRUCache struct {
	cap   int
	cache map[int]*Node
	head  *Node // dummy head
	tail  *Node // dummy tail
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (c *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *LRUCache) addToFront(node *Node) {
	node.Next = c.head.Next
	node.Prev = c.head

	c.head.Next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) moveToFront(node *Node) {
	c.remove(node)
	c.addToFront(node)
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.Value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		node.Value = value
		c.moveToFront(node)
		return
	}

	node := &Node{Key: key, Value: value}
	c.cache[key] = node
	c.addToFront(node)

	if len(c.cache) > c.cap {
		lru := c.tail.Prev
		c.remove(lru)
		delete(c.cache, lru.Key)
	}
}
