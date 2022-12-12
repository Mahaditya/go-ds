# Data Structures in Go

The goal of this project is to implement all major data structures in go to facilitate development. 


### Vector

```go
package main

import (
	"fmt"
	"go-ds/src/dstructs/vector"
)

func main() {
	var arr vector.Vector[string]
	arr.Push("alpha").Push("beta")
	fmt.Println(arr)
}

> Output: <Vector [alpha beta]>
```


### Stack 

```go
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

> Output: 

Stack [alpha beta gamma]<---
Stack [alpha beta]<---
beta

```
