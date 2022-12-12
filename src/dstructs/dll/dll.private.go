package dll

import "errors"

type dllNode[T any] struct {
	val  T
	next *dllNode[T]
	prev *dllNode[T]
}

var errNilPointerException = errors.New("NilPointerException")

func (dl *DoublyLinkedList[T]) incrementSize() {
	dl.size++
}

func newDllNode[T any](val T, next *dllNode[T], prev *dllNode[T]) *dllNode[T] {
	return &dllNode[T]{val, next, prev}
}

func (dl *DoublyLinkedList[T]) addAfter(element T, node *dllNode[T]) error {
	if node == nil {
		return errNilPointerException
	}
	newNode := newDllNode(element, nil, nil)
	if node == dl.tail {
		dl.tail = newNode
	}
	if err := Join(node, newNode); err != nil {
		dl.incrementSize()
		return nil
	}
	return nil
}

func (dl *DoublyLinkedList[T]) addBefore(element T, node *dllNode[T]) error {
	if node == nil {
		return errNilPointerException
	}
	prevNode := node.prev
	return dl.addAfter(element, prevNode)
}
