package bst

import (
	"fmt"
	"golang.org/x/exp/constraints"
	
)

type BinarySearchTree[K constraints.Ordered, T any] struct {
	root *bstNode[K, T]
	size uint64
}

func (bst *BinarySearchTree[K, T]) Create(key K, val T) *BinarySearchTree[K, T] {
	return bst.Insert(key,val)
}

func (bst *BinarySearchTree[K, T]) Insert(key K, val T) *BinarySearchTree[K, T] {
	if bst.root == nil {
		bst.root = &bstNode[K, T]{nil, nil, key, val}
		bst.size++
		return bst
	}
	_, lowerLimit := bst.search(key)

	if lowerLimit.key < key {
		lowerLimit.right = &bstNode[K, T]{nil, nil, key, val}
		bst.size++
	} else {
		lowerLimit.left = &bstNode[K, T]{nil, nil, key, val}
		bst.size++
	}
	return bst
}



func (bst *BinarySearchTree[K, T]) Search(searchKey K) (found bool, val T) {
	found, node := bst.search(searchKey)
	if node != nil {
		return found, node.value
	}
	return found, node.value
}

func (bst *BinarySearchTree[K, T]) Size() uint64 {
	return bst.size
}

func (bst *BinarySearchTree[K, T]) InOrderKeys() []K {
	inOrderKeysContainer := make([]K, 0)
	inOrderNodes := bst.inOrder()
	for _, node := range inOrderNodes {
		inOrderKeysContainer = append(inOrderKeysContainer, node.key)
	}
	return inOrderKeysContainer
}

func (bst *BinarySearchTree[K, T]) InOrderValues() []T {
	inOrderValuesContainer := make([]T, 0)
	inOrderNodes := bst.inOrder()
	for _, node := range inOrderNodes {
		inOrderValuesContainer = append(inOrderValuesContainer, node.value)
	}
	return inOrderValuesContainer
}


func (bst *BinarySearchTree[K, T]) Print() {
	levelOrder:= bst.levelOrderTraversal()
	for _,level := range levelOrder{
		for _,node:= range level{
			fmt.Print(node.key," ")
		}
		fmt.Println()
	}
}