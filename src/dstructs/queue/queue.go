package queue

import (
	"fmt"
	"go-ds/src/dstructs/vector"
)

type Queue[T any] struct {
	data vector.Vector[T]
}

func (q *Queue[T]) Size() int {
	return q.data.Size()
}

func (q *Queue[T]) Push(element T) {
	q.data.Push(element)
}

func (q *Queue[T]) Front() (T, error) {
	return q.data.Front()
}

func (q *Queue[T]) Pop() (T, error) {
	return q.data.PopFront()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue[T]) IsNotEmpty() bool {
	return q.data.IsNotEmpty()
}

func (q Queue[T]) String() string {
	return fmt.Sprintf("Queue --->%s", q.data.ToString())
}
