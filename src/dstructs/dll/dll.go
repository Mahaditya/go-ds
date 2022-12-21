package dll

import "fmt"

type DoublyLinkedList[T any] struct {
	head *dllNode[T]
	tail *dllNode[T]
	size int
}

func (dl *DoublyLinkedList[T]) GetHead() *dllNode[T] {
	return dl.head
}

func (dl *DoublyLinkedList[T]) GetTail() *dllNode[T] {
	return dl.tail
}

func (dl *DoublyLinkedList[T]) AddAfter(element T, node *dllNode[T]) error {
	return dl.addAfter(element, node)
}

// func (dl *DoublyLinkedList[T]) AddBefore(element T, node *dllNode[T]) error {
// 	return dl.addBefore(element, node)
// }



func (dl *DoublyLinkedList[T]) AddLast(element T) error {
	return dl.addAfter(element, dl.GetTail())
}

func (dl *DoublyLinkedList[T]) AddFirst(element T) error {
	return dl.addBefore(element, dl.GetHead())
}

func (sl *DoublyLinkedList[T]) ForEach(f func(node *dllNode[T])) {
	currNode := sl.head
	for currNode != nil {
		f(currNode)
		currNode = currNode.next
	}
}

func (sl *DoublyLinkedList[T]) ForEachReverse(f func(node *dllNode[T])) {
	currNode := sl.tail
	for currNode != nil {
		f(currNode)
		currNode = currNode.prev
	}
}

func (dl DoublyLinkedList[T]) Format(f fmt.State, verb rune) {
	if dl.size == 0 {
		fmt.Print("{empty}")
		return
	}
	dl.ForEach(func(node *dllNode[T]) {
		if node.next != nil && node.prev != nil {
			fmt.Printf(" |%v| <=>", node.val)
		} else if node.next == nil && node.prev == nil {
			fmt.Printf("nil<- |%v| ->nil", node.val)
		} else if node.next == nil {
			fmt.Printf(" |%v| ->nil", node.val)
		} else {
			fmt.Printf("nil<- |%v| <=>", node.val)
		}
	})
}
