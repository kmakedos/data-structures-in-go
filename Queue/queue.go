package Queue

import (
	"fmt"
)

type Node struct {
	Data int
	next *Node
	prev *Node
}

type Queue struct {
	head *Node
	tail *Node
	len  int64
}

func (n *Node) Create(data int) *Node {
	n.Data = data
	n.next = nil
	n.prev = nil
	return n
}

func (q *Queue) Create() *Queue {
	q = new(Queue)
	q.head = nil
	q.tail = nil
	return q
}

func (q *Queue) Enqueue(n *Node) {
	if q.head != nil {
		q.head.prev = n
		n.next = q.head
		q.head = n
	} else {
		q.head = n
		q.tail = n
	}
}

func (q *Queue) Dequeue() *Node {
	n := q.tail
	if q.tail != nil {
		if q.tail.prev != nil {
			q.tail = q.tail.prev
			q.tail.next = nil
		} else {
			q.head = nil
			q.tail = nil
		}
	}
	return n
}

func (q *Queue) String() string {
	var s string
	for i := q.head; i != nil; i = i.next {
		s += fmt.Sprintf("%d ", i.Data)
	}
	return s
}
