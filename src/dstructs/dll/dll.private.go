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

func (dl *DoublyLinkedList[T]) decrementSize() {
	dl.size--
}

func newDllNode[T any](val T, next *dllNode[T], prev *dllNode[T]) *dllNode[T] {
	return &dllNode[T]{val, next, prev}
}

func (dl *DoublyLinkedList[T]) addAfterIndex(element T, index int){
	
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


func (node *dllNode[T]) addNext(element T) {
	node.next = &dllNode[T]{element,nil,node}
}

func (node *dllNode[T]) addPrev(element T) {
	node.prev = &dllNode[T]{element,node,nil}
}