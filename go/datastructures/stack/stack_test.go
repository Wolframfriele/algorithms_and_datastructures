package stack

import (
	"testing"
)

func TestStackEmpty(t *testing.T) {
	stack := new(Stack)
	_, err := stack.pop()
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.append(1)
	stack.append(2)
	got, err := stack.pop()
	want1 := 2
	if got != want1 {
		t.Errorf("Got %v want %v", got, want1)
	}
	got, err = stack.pop()
	want2 := 1
	if got != want2 {
		t.Errorf("Got %v want %v", got, want2)
	}
	got, err = stack.pop()
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
