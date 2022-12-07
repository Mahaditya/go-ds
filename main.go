package main

import (
	"fmt"
	"go-ds/src/dstructs/queue"
	"go-ds/src/dstructs/stack"
)

func main(){
	q:= new(queue.Queue[int])
	q.Push(1).Push(2).Push(3)
	
	for q.IsNotEmpty(){
		v,_:=q.Pop()
		fmt.Println(v)
	}
	st:= new(stack.Stack[int])
	st.Push(1).Push(2).Push(3).Push(4)
	for st.IsNotEmpty(){
		val,_:=st.Pop()
		fmt.Println(val)
	}
}