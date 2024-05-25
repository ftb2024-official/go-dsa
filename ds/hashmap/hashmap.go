package hashmap

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   any
	value any
	next  *node
}

type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

// New returns a new instance of HashMap
func New() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

// Make returns a new instance of HashMap with input size and capacity
func Make(size, capacity uint64) HashMap {
	return HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, defaultCapacity),
	}
}

func (hm *HashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}
