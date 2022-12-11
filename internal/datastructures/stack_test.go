package datastructures

import "testing"

func  TestNewS(t *testing.T) {
	// GIVEN
	// WHEN
	stack := NewStack()
	// THEN
	if stack == nil {
		t.Error("NewStack() returned nil")
	}
}

func TestPush(t *testing.T) {
	// GIVEN
	stack := NewStack()
	// WHEN
	stack.Push(1)
	// THEN
	if stack.Len() != 1 {
		t.Errorf("Push() did not add element to stack")
	}
}

func TestPop(t *testing.T)
	// GIVEN
	stack := NewStack()
	stack.Push(1)
	// WHEN
	val := stack.Pop()
	// THEN
	if val != 1 {
		t.Errorf("Pop() did not return correct value")
	}
}