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