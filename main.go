package main

import (
	"fmt"
	"go-ds/src/dstructs"
	"go-ds/src/dstructs/vector"
)

func call(v dstructs.Iterable[int]){
	it:= v.GetIterator()
	for it.HasNext() {
		fmt.Println(it.GetNext())
	}
}


func main() {
	var v vector.Vector[int]
	v.Push(1)
	v.Push(2)
	v.Push(3)
	call(&v)
}
