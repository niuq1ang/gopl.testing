package queue

import (
	"sync"
)

type Queue struct {
	firstElement *Element
	lastElement  *Element
	size         int64
	lock         sync.Mutex
}

type Element struct {
	next *Element
	data []byte
}

func New() *Queue {
	return new(Queue)
}

func (q *Queue) Put(data []byte) {
	q.lock.Lock()
	defer q.lock.Unlock()

	node := &Element{
		data: data,
	}
	if q.lastElement == nil {
		q.firstElement = node
	} else {
		q.lastElement.next = node
	}
	q.lastElement = node
	q.size++
}
