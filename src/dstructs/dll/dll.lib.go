package dll

import "errors"

func Make[T any](element T) *DoublyLinkedList[T] {
	var dl DoublyLinkedList[T]
	dl.head = newDllNode(element, nil, nil)
	dl.tail = dl.head
	dl.incrementSize()
	return &dl
}

func Join[T any](node1 *dllNode[T], node2 *dllNode[T]) error {
	if node1 == nil || node2 == nil {
		return errors.New("nil node passed")
	}
	node1.next = node2
	node2.prev = node1
	return nil
}
