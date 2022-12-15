package sll

import (
	"errors"
	"fmt"
)

type sllIterator[T any] struct {
	index       int
	currentNode *lLNode[T]
	sll         *SinglyLinkedList[T]
}

func (sli *sllIterator[T]) HasNext() bool {
	return sli.index < sli.sll.Size()
}

func (sli *sllIterator[T]) GetNext() (T, error) {
	if sli.HasNext() {
		if currVal, err := getVal(sli.currentNode); err != nil {
			return *new(T), err
		} else {
			sli.index++
			sli.currentNode, _ = getNext(sli.currentNode)
			return currVal, err
		}
	}
	fmt.Println("Doesn't have the elements")
	return *new(T), errors.New("no next element")
}
