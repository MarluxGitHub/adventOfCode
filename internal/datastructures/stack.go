package datastructures

import "fmt"

type stack struct {
	data []interface{}
}

func NewStack() *stack {
	return &stack{make([]interface{}, 0)}
}

func (s *stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *stack) Peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) Clear() {
	s.data = make([]interface{}, 0)
}

func (s *stack) Values() []interface{} {
	return s.data
}

func (s *stack) String() string {
	return fmt.Sprint(s.data)
}