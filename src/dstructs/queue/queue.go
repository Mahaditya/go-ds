package queue

import (
	"fmt"
	"go-ds/src/dstructs/vector"
)

type Queue[T any] struct{
	data vector.Vector[T]
}

func (q *Queue[T]) Size() int{
	return q.data.Size()
}

func (q *Queue[T]) Push(element T) *Queue[T]{
	q.data.Push(element)
	return q
}

func (q *Queue[T]) Front() (T,error){
	return q.data.First()
}

func (q *Queue[T]) Pop() (T,error){
	return q.data.PopFront()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue[T]) IsNotEmpty() bool {
	return q.data.IsNotEmpty()
}

func (q *Queue[T]) Print()  {
	fmt.Print("->")
	q.data.Print()
}