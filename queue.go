package graph

type queueItem struct {
	value interface{}
	next  *queueItem
}

type queue struct {
	front, back *queueItem
	len         int
}

func (q *queue) enqueue(i interface{}) {
	n := &queueItem{
		value: i,
	}
	if q.len == 0 {
		q.front = n
		q.back = n
	} else {
		q.back.next = n
		q.back = n
	}
	q.len++
}

func (q *queue) dequeue() interface{} {
	if q.len == 0 {
		return nil
	}
	n := q.front
	if q.len == 1 {
		q.front = nil
		q.back = nil
	} else {
		q.front = q.front.next
	}
	q.len--

	return n.value
}
