package gocollections

type IQueue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
}

type IntQueue struct {
	size  int
	queue []int
}

func (q *IntQueue) Enqueue(e interface{}) {
	q.queue = append(q.queue, e.(int))
	q.size += 1
}

func (q *IntQueue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	q.size -= 1
	return val
}

func (q *IntQueue) Peek() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.queue[0]
}

func (q *IntQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *IntQueue) Size() int {
	return q.size
}
