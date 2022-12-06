package queue

import (
	"errors"
	"fmt"
)

type Queue[T any] struct{
	data []T
}

func (q *Queue[T]) Size() int{
	return len(q.data)
}

func (q *Queue[T]) Push(data T){
	q.data = append(q.data,data)
}

func (q *Queue[T]) Front() (T,error){
	if(q.Size()>0){
		return q.data[0],nil
	}
	return *new(T),errors.New("queue is empty") 
}

func (q *Queue[T]) Pop() (error){
	if(q.Size()>0){
		q.data=q.data[1:]
		return nil
	}
	return errors.New("popping an empty queue")
}

func (q *Queue[T]) NotEmpty() bool {
	return q.Size()>0
}

func (q *Queue[T]) Print()  {
	fmt.Println("->",q.data)
}