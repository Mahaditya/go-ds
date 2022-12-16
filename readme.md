# Data Structures in Go

The goal of this project is to implement all major data structures in go to facilitate development. 

## Major Features:

- Complete Generic Support 🚀
- All major Data Structures ✅
- Supports Polymorphism through Interfaces 🟨 🔷
- Pretty Printing by default for better visualization `<Vector [1 2]>`



#### How to import

```go
import (
	"go-ds/src/dstructs/vector"
	"go-ds/src/dstructs/stack"
)

```

### Vector

```go
var vect vector.Vector[string]
vect.Push("alpha")
vect.Push("beta")
fmt.Println(arr) // <Vector [alpha beta]>

var vectInt vector.Vector[int]

vectInt.Push(1)
vectInt.push(99)

fmt.Println(vectInt) // <Vector [1 99]>


```


### Stack 

```go
var myStack stack.Stack[string]

myStack.Push("alpha")
myStack.Push("beta")
myStack.Push("gamma")

fmt.Println(myStack) // Stack [alpha beta gamma]<---

myStack.Pop()

fmt.Println(myStack) // Stack [alpha beta]<---

if top,err:= myStack.Top(); err!=nil{
  fmt.Println(err)
}else{
  fmt.Println(top) // beta
}

```

### Pair

```go

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

```