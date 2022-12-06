package main

import (
	"go-ds/src/dstructs/queue"
)

func main(){
	q:= new(queue.Queue[int])
	q.Push(1).Push(2).Push(3)
	
	for q.IsNotEmpty(){
		q.Print()
		q.Pop()
	}
}