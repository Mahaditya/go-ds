package sll

import (
	"errors"
	"fmt"
)

type SinglyLinkedList[T any] struct {
	head *lLNode[T]
	tail *lLNode[T]
	size uint64
}

func (sl *SinglyLinkedList[T]) GetHead() *lLNode[T] {
	return sl.head
}

func (sl *SinglyLinkedList[T]) GetTail() *lLNode[T] {
	return sl.tail
}

func (sl *SinglyLinkedList[T]) AddFirst(element T) {
	sl.addBefore(element, sl.head)
}

func (sl *SinglyLinkedList[T]) AddLast(element T) {
	sl.addAfter(element, sl.tail)
}

func (sl *SinglyLinkedList[T]) AddAfter(node *lLNode[T], element T) error {
	return sl.addAfter(element, node)
}

func (sl *SinglyLinkedList[T]) AddBefore(node *lLNode[T], element T) error {
	return sl.addBefore(element, node)
}

func (sl *SinglyLinkedList[T]) ForEach(f func(node *lLNode[T])) {
	currNode := sl.head
	for currNode != nil {
		f(currNode)
		currNode = currNode.next
	}
}

func (sl *SinglyLinkedList[T]) Remove(node *lLNode[T]) error {
	prevNode := sl.getPrev(node)
	if prevNode != nil {
		prevNode.next = node.next
		return nil
	}
	if node == sl.head {
		sl.head = node.next
		return nil
	}
	return errors.New("node not found")
}

func (sl *SinglyLinkedList[T]) Pop(node *lLNode[T]) error {
	return sl.Remove(sl.GetTail())
}

func (sl *SinglyLinkedList[T]) PopFront(node *lLNode[T]) error {
	return sl.Remove(sl.GetHead())
}

func (sl SinglyLinkedList[T]) Format(f fmt.State, verb rune) {
	if sl.size == 0 {
		fmt.Print("{empty}")
		return
	}
	sl.ForEach(func(node *lLNode[T]) {
		if node.next != nil {
			fmt.Print(node.val, "->")
		} else {
			fmt.Print(node.val, "->nil")
		}
	})
}
