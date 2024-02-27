package queue

import (
	"errors"
)

type Node struct {
	next  *Node
	value int
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) append(val int) {
	if q.length == 0 {
		node := Node{next: nil, value: val}
		q.head = &node
		q.tail = &node
	} else {
		node := Node{next: nil, value: val}
		q.tail.next = &node
		q.tail = &node
	}
	q.length += 1
}

func (q *Queue) pop() (int, error) {
	if q.length == 0 {
		return 0, errors.New("No items in queue")
	} else {
		node := q.head
		q.head = node.next
		q.length -= 1
		return node.value, nil
	}
}

func (q *Queue) peek() (int, error) {
	if q.length == 0 {
		return 0, errors.New("No items in queue")
	} else {
		return q.head.value, nil
	}
}
