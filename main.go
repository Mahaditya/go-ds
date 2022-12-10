package main

import (
	"fmt"
	"go-ds/src/dstructs/sll"
)

func main() {
	var linkedList sll.SinglyLinkedList[int]
	linkedList.AddLast(9) 
	linkedList.AddFirst(1)
	linkedList.AddFirst(2)
	linkedList.AddFirst(6)
	linkedList.AddLast(8)
	head:= linkedList.GetHead()
	linkedList.AddAfter(19,head)
	fmt.Println(linkedList)
	fmt.Println(linkedList.GetHead(),linkedList.GetTail())
}
