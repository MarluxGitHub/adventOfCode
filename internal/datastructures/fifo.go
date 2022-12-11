package datastructures

import "fmt"

type fifo struct {
	data []interface{}
}

func NewFIFO() *fifo {
	return &fifo{make([]interface{}, 0)}
}

func (f *fifo) Push(v interface{}) {
	f.data = append(f.data, v)
}

func (f *fifo) Pop() interface{} {
	if len(f.data) == 0 {
		return nil
	}

	v := f.data[0]
	f.data = f.data[1:]
	return v
}

func (f *fifo) Peek() interface{} {
	if len(f.data) == 0 {
		return nil
	}

	return f.data[0]
}

func (f *fifo) Len() int {
	return len(f.data)
}

func (f *fifo) IsEmpty() bool {
	return len(f.data) == 0
}

func (f *fifo) Clear() {
	f.data = make([]interface{}, 0)
}

func (f *fifo) Values() []interface{} {
	return f.data
}

func (f *fifo) String() string {
	return fmt.Sprint(f.data)
}