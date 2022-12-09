package main

import (
	"fmt"
	"go-ds/src/dstructs/stack"
)

func main() {
	var st stack.Stack[int]
	if a,err:= st.Pop();err!=nil{
		fmt.Println(a)
	}else{
		fmt.Println(a)
	}
}
