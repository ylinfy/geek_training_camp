// 使用双端链表和无序map实现LRU cache
type LRUCache struct {
	size       int                  // 使用了多少空间
	capacity   int                  // 缓存空间
	cache      map[int]*DLinkedNode // 无序map, 方便O(1)查找key对应的缓存节点
	head, tail *DLinkedNode         // 缓存节点全部连接在head和tail中间，形成一个双链表
}

type DLinkedNode struct {
	key        int // key对应cache中的key, 在删除节点时，方便使用delete cache[node.key]直接删除map中对应的key
	value      int // 当前缓存节点的值
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key, value: value}
}

// 初始化缓存
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// 获取缓存值
func (L *LRUCache) Get(key int) int {
	node, ok := L.cache[key]
	if !ok {
		return -1
	}

	L.moveToHead(node)
	return node.value
}

// 插入缓存值
func (L *LRUCache) Put(key, value int) {
	if _, ok := L.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		L.cache[key] = node
		// 没有使用过该值，缓存在最近端，即链表最前面的位置
		L.addToHead(node)
		L.size++
		// 当缓存数据超过capacity时，删除最末尾的node，即比较旧的数据，而保留最近使用的值
		if L.size > L.capacity {
			removed := L.removeTail()
			delete(L.cache, removed.key)
			L.size--
		}
	} else {
		// 最近使用的且已经缓存过，直接移到最前面
		node := L.cache[key]
		node.value = value
		L.moveToHead(node)
	}
}

func (L *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = L.head
	node.next = L.head.next
	L.head.next.prev = node
	L.head.next = node
}

func (L *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (L *LRUCache) moveToHead(node *DLinkedNode) {
	L.removeNode(node)
	L.addToHead(node)
}

func (L *LRUCache) removeTail() *DLinkedNode {
	node := L.tail.prev
	L.removeNode(node)
	return node
}
