package lua

import (
	"sync"
)

const (
	DEFAULT_CACHE_SIZE = 16
)

type Queue struct {
	list     []interface{}
	rIdx     uint32
	wIdx     uint32
	capacity uint32
	mask     uint32
	max      uint32
}

func NewQueue() *Queue {
	return &Queue{}
}

func roundup_pow_of_two(size uint32) (uint32, uint32) {
	log2 := 0
	if size == 0 {
		return 0, 0
	}
	if size & 0xFFFF0000 {
		log2 += 16
		size >>= 16
	}
	if size & 0xFF00 {
		log2 += 8
		size >>= 8
	}
	if size & 0xF0 {
		log2 += 4
		size >>= 4
	}
	if size & 0xC {
		log2 += 2
		size >>= 2
	}
	if size & 0x2 {
		log2 += 1
		size >>= 1
	}

	if (size & (size - 1)) != 0 {
		log2 += 1
	}
	return (1 << log2), log2
}

func (q *Queue) Init(max uint32) {
	size, _ := roundup_pow_of_two(max)
	q.capacity = size
	q.max = max
	q.mask = size - 1
	q.rIdx = 0
	q.wIdx = 0
}

func (q *Queue) Get() interface{} {
	if q.wIdx-q.rIdx <= 0 {
		return nil
	}
	node := q.list[q.rIdx&q.mask]
	q.rIdx += 1

	return node
}

func (q *Queue) GetMulti(count uint32, dst *[]interface{}) uint32 {
	used := q.wIdx - q.rIdx
	if (used <= 0) || (dst == nil) {
		return 0
	}
	if used < count {
		count = used
	}

	rIdx := q.rIdx & q.mask
	left := count - (q.capacity - rIdx)
	i := 0
	if left <= 0 {
		for i = 0; i < count; i++ {
			dst[i] = q.list[rIdx+i]
		}
	} else {
		copyed := q.capacity - rIdx
		for i = 0; i < copyed; i++ {
			dst[i] = q.list[rIdx+i]
		}
		for i = 0; i < left; i++ {
			dst[copyed+i] = q.list[i]
		}
	}
	q.rIdx += count

	return count
}

func (q *Queue) Put(v interface{}) bool {
	if q.wIdx-q.rIdx == q.max {
		return false
	}
	q.list[q.wIdx&q.mask] = v
	q.wIdx += 1

	return true
}

func (q *Queue) PutMulti(count uint32, v []interface{}) uint32 {
	used := q.wIdx - q.rIdx
	if used == q.max {
		return 0
	}
	idle := q.max - used
	if idle < count {
		count = idle
	}

	wIdx := q.wIdx & q.mask
	left := count - (q.capacity - wIdx)
	if left <= 0 {
		for i := 0; i < count; i++ {
			q.list[wIdx+i] = v[i]
		}
	} else {
		for i := 0; i < q.capacity-wIdx; i++ {
			q.list[wIdx+i] = v[i]
		}
		base := i
		for i := 0; i < left; i++ {
			q.list[i] = v[base+i]
		}
	}
	q.wIdx += count

	return count
}

type Pool struct {
	cache       []*Queue
	concurrency uint32
	cacheSize   uint32

	backup   *Queue
	lock     sync.Mutex
	maxCount uint32
	curCount uint32
	op       *NodeOperation
}

type NodeOperation interface {
	New() interface{}
}

func NewPool(maxCount uint32, cacheSize uint32, concurrency uint32) *Pool {
	pool := &Pool{maxCount: maxCount, cacheSize: cacheSize, concurrency: concurrency}
	pool.cache = make([]*Queue, 0, concurrency)
	for i := 0; i < concurrency; i++ {
		pool.cache[i] = NewQueue()
		pool.cache[i].Init(DEFAULT_CACHE_SIZE)
	}

	return pool
}

func (p *Pool) SetNodeOperation(op *NodeOperation) {
	p.op = op
}

func (p *Pool) FillUpNode() {
	for _, queue := range p.cache {
		for i := 0; i < queue.max; i++ {
			queue.Put(p.op.New())
		}
	}
}

func (p *Pool) GetNode(idx uint32) interface{} {
	queue := p.cache[idx]
	node := queue.Get()

	return node
}

func (p *Pool) PutNode(idx uint32, data interface{}) bool {
	queue := p.cache[idx]
	res := queue.Put(data)

	return res
}
