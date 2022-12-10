package sll

import "errors"

type lLNode[T any] struct {
	val  T
	next *lLNode[T]
}

func getNilPointerException() error {
	return errors.New("NilPointerException")
}

func newLLNode[T any](val T, next *lLNode[T]) *lLNode[T] {
	return &lLNode[T]{val, next}
}

func (sl *SinglyLinkedList[T]) incrementSize() {
	sl.size++
}

func setNext[T any](node *lLNode[T], nextNode *lLNode[T]) error {
	if node != nil {
		node.next = nextNode
		return nil
	}
	return getNilPointerException()
}

func setVal[T any](node *lLNode[T], value T) error {
	if node != nil {
		node.val = value
		return nil
	}
	return getNilPointerException()
}

func getVal[T any](node *lLNode[T]) (T, error) {
	if node != nil {
		return node.val, nil
	}
	return *new(T), getNilPointerException()
}

func getNext[T any](node *lLNode[T]) (*lLNode[T], error) {
	if node != nil {
		return node.next, nil
	}
	return nil, getNilPointerException()
}

func (sl *SinglyLinkedList[T]) getPrev(currNode *lLNode[T]) *lLNode[T] {
	head := sl.GetHead()
	if head == currNode {
		return nil
	}
	for head != nil {
		if head.next == currNode {
			return head
		}
		head = head.next
	}
	return nil
}

func (sl *SinglyLinkedList[T]) addAfter(val T, node *lLNode[T]) error {
	nextNode, err := getNext(node)
	newNode := newLLNode(val, nextNode)
	if err != nil{
		sl.head = newNode
	}
	setNext(node, newNode)
	if node == sl.tail {
		sl.tail = newNode
	}
	sl.incrementSize()
	return nil
}

func (sl *SinglyLinkedList[T]) addBefore(val T, node *lLNode[T]) error {
	prevNode := sl.getPrev(node)
	if prevNode == nil {
		newNode := newLLNode(val, node)
		sl.head = newNode
		if sl.tail==nil{
			sl.tail = sl.head
		}
		sl.incrementSize()
		return nil
	}
	return sl.addAfter(val, prevNode)
}
