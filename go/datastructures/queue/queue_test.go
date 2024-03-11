package queue

import (
	"testing"
)

func TestQueueEmpty(t *testing.T) {
	queue := new(Queue)
	_, err := queue.deque()
	if err == nil {
		t.Errorf("Expected empty queue error")
	}
}

func TestQueueOnce(t *testing.T) {
	queue := new(Queue)
	queue.enqueue(1)
	got, err := queue.deque()
	if got != 1 || err != nil {
		t.Errorf("Got %v want %v", got, 1)
	}
	got, err = queue.deque()
	if err == nil {
		t.Errorf("Expected empty queue error")
	}
}

func TestQueueTwice(t *testing.T) {
	queue := new(Queue)
	queue.enqueue(1)
	queue.enqueue(2)
	got, err := queue.deque()
	if got != 1 || err != nil {
		t.Errorf("Got %v want %v", got, 1)
	}
	got, err = queue.deque()
	if got != 2 || err != nil {
		t.Errorf("Got %v want %v", got, 2)
	}
	got, err = queue.deque()
	if err == nil {
		t.Errorf("Expected empty queue error")
	}
}

func TestQueueThrice(t *testing.T) {
	queue := new(Queue)
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	got, err := queue.deque()
	if got != 1 || err != nil {
		t.Errorf("Got %v want %v", got, 1)
	}
	got, err = queue.deque()
	if got != 2 || err != nil {
		t.Errorf("Got %v want %v", got, 2)
	}
	got, err = queue.deque()
	if got != 3 || err != nil {
		t.Errorf("Got %v want %v", got, 3)
	}
	got, err = queue.deque()
	if err == nil {
		t.Errorf("Expected empty queue error")
	}
}

func TestQueuePeekEmpty(t *testing.T) {
	queue := new(Queue)
	_, err := queue.peek()
	if err == nil {
		t.Errorf("Expected empty queue error")
	}
}

func TestQueuePeek(t *testing.T) {
	queue := new(Queue)
	queue.enqueue(1)
	got, err := queue.peek()
	if got != 1 || err != nil {
		t.Errorf("Got %v want %v", got, 1)
	}
}
