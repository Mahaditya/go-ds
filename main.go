package main

import (
	"fmt"
	"go-ds/src/dstructs/stack"
)

func main() {
	var myStack stack.Stack[string]

	myStack.Push("alpha")
	myStack.Push("beta")
	myStack.Push("gamma")

	fmt.Println(myStack)

	myStack.Pop()

	fmt.Println(myStack)

	if top,err:= myStack.Top(); err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(top)
	}
	
}
