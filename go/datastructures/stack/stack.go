package stack

import (
	"errors"
)

type Node struct {
	previous *Node
	value    int
}

type Stack struct {
	length int
	head   *Node
	tail   *Node
}

func (s *Stack) append(val int) {
	if s.length == 0 {
		node := Node{value: val}
		s.head = &node
		s.tail = &node
		s.head.previous = s.tail
	} else {
		node := Node{value: val, previous: s.head}
		s.head = &node
	}
	s.length += 1
}

func (s *Stack) pop() (int, error) {
	if s.length > 0 {
		lastNode := s.head
		s.head = lastNode.previous
		s.length -= 1
		return lastNode.value, nil
	} else {
		return -1, errors.New("No items on stack.")
	}
}
