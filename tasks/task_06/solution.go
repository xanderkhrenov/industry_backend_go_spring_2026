package main

type LRUCache[K comparable, V any] struct {
	storage  map[K]*node[K, V]
	capacity int
	head     *node[K, V]
	tail     *node[K, V]
}

type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		storage:  make(map[K]*node[K, V], capacity),
		capacity: capacity,
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	curNode, ok := c.storage[key]
	if c.capacity == 0 || !ok {
		return value, false
	}
	c.moveToHead(curNode)
	return curNode.value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity == 0 {
		return
	}

	curNode, ok := c.storage[key]
	if ok {
		curNode.value = value
		c.moveToHead(curNode)
		return
	}

	c.addToHead(&node[K, V]{key: key, value: value})
	if c.capacity < len(c.storage) {
		c.deleteFromTail()
	}
}

func (c *LRUCache[K, V]) moveToHead(curNode *node[K, V]) {
	if c.head == curNode {
		return
	}

	prev, next := curNode.prev, curNode.next
	curNode.prev = nil
	prev.next = next
	if c.tail == curNode {
		c.tail = prev
	} else {
		curNode.next = nil
		next.prev = prev
	}
	c.addToHead(curNode)
}

func (c *LRUCache[K, V]) addToHead(curNode *node[K, V]) {
	c.storage[curNode.key] = curNode
	if len(c.storage) == 1 {
		c.head = curNode
		c.tail = curNode
		return
	}
	curNode.next = c.head
	c.head.prev = curNode
	c.head = curNode
}

func (c *LRUCache[K, V]) deleteFromTail() {
	delete(c.storage, c.tail.key)
	c.tail = c.tail.prev
	c.tail.next = nil
}
