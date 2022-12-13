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
	arr.Push("alpha")
	arr.Push("beta")
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

### Pair

```go

import (
	"fmt"
	"go-ds/src/dstructs/pair"
	"go-ds/src/dstructs/stack"
)

func main() {
	catPair:= pair.Make("Cat",4)
	fmt.Println(catPair) //> {Cat 4}

	birdPair:= pair.Make("Bird",2)
	fmt.Println(birdPair) //> {Bird 2}

	var myStack stack.Stack[pair.Pair[string,int]]
	myStack.Push(catPair)
	myStack.Push(birdPair)
	fmt.Println(myStack)  //> Stack [{Cat 4} {Bird 2}]<---

	if topPair,err:= myStack.Top(); err==nil{
		fmt.Println("TopPair",topPair) //> TopPair {Bird 2}
	}
}

```