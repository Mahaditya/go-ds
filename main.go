package main

import (
	"fmt"
	"go-ds/src/dstructs/dll"
)

func main() {
	linkedList:=dll.Make(9)
	linkedList.AddLast(8)
	linkedList.AddFirst(7)
	linkedList.AddLast(1)
	fmt.Println(linkedList)
}
