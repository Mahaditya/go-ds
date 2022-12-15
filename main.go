package main

import (
	"fmt"
	"go-ds/src/dstructs"
	"go-ds/src/dstructs/sll"
)

func call(v dstructs.Iterable[int]) {
	it := v.GetIterator()
	for it.HasNext() {
		fmt.Println(it.GetNext())
	}
}

func main() {
	sl := sll.Make(1)
	sl.AddLast(7)
	sl.AddFirst(5)
	call(sl)
}
