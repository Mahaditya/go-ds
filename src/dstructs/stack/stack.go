package stack

import "go-ds/src/dstructs/vector"

type Stack[T any] struct {
	data vector.Vector[T]
}

func (st *Stack[T]) Push(element T) *Stack[T] {
	st.data.Push(element)
	return st
}

func (st *Stack[T]) Pop() (T, error) {
	return st.data.Pop()
}

func (st *Stack[T]) Top() (T, error) {
	return st.data.First()
}

func (st *Stack[T]) Size() int {
	return st.data.Size()
}
func (st *Stack[T]) IsEmpty() bool {
	return st.data.Size()<=0
}
func (st *Stack[T]) IsNotEmpty() bool {
	return !st.IsEmpty()
}
