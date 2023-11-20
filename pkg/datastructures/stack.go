package datastructures

import "fmt"

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0)}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) Peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Clear() {
	s.data = make([]interface{}, 0)
}

func (s *Stack) Values() []interface{} {
	return s.data
}

func (s *Stack) String() string {
	return fmt.Sprint(s.data)
}