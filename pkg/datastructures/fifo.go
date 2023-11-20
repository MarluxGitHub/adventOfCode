package datastructures

import "fmt"

type Fifo struct {
	data []interface{}
}

func NewFIFO() *Fifo {
	return &Fifo{make([]interface{}, 0)}
}

func (f *Fifo) Push(v interface{}) {
	f.data = append(f.data, v)
}

func (f *Fifo) Pop() interface{} {
	if len(f.data) == 0 {
		return nil
	}

	v := f.data[0]
	f.data = f.data[1:]
	return v
}

func (f *Fifo) Peek() interface{} {
	if len(f.data) == 0 {
		return nil
	}

	return f.data[0]
}

func (f *Fifo) Len() int {
	return len(f.data)
}

func (f *Fifo) IsEmpty() bool {
	return len(f.data) == 0
}

func (f *Fifo) Clear() {
	f.data = make([]interface{}, 0)
}

func (f *Fifo) Values() []interface{} {
	return f.data
}

func (f *Fifo) String() string {
	return fmt.Sprint(f.data)
}
