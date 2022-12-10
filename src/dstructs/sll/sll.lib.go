package sll

func Make[T any](element T) *SinglyLinkedList[T]{
	var sl SinglyLinkedList[T]
	sl.head = newLLNode(element, nil)
	sl.tail = sl.head
	sl.size++
	return &sl
}