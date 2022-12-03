package dstructs

import (
	"golang.org/x/exp/constraints"
)

type binarySearchTreeNode[K constraints.Ordered, T any] struct {
	left  *binarySearchTreeNode[K, T]
	right *binarySearchTreeNode[K, T]
	key   K
	value T
}

type BinarySearchTree[K constraints.Ordered, T any] struct {
	root *binarySearchTreeNode[K, T]
	size uint64
}

func (bst *BinarySearchTree[K, T]) Create(key K, val T) *BinarySearchTree[K, T] {
	return bst.Insert(key,val)
}

func (bst *BinarySearchTree[K, T]) Insert(key K, val T) *BinarySearchTree[K, T] {
	if bst.root == nil {
		bst.root = &binarySearchTreeNode[K, T]{nil, nil, key, val}
		bst.size++
		return bst
	}
	_, lowerLimit := bst.search(key)

	if lowerLimit.key < key {
		lowerLimit.right = &binarySearchTreeNode[K, T]{nil, nil, key, val}
		bst.size++
	} else {
		lowerLimit.left = &binarySearchTreeNode[K, T]{nil, nil, key, val}
		bst.size++
	}
	return bst
}

func (bst *BinarySearchTree[K, T]) search(searchKey K) (found bool, value *binarySearchTreeNode[K, T]) {
	if bst.root == nil {
		return false, bst.root
	}
	node := bst.root
	for node != nil{
		if node.key > searchKey && node.left != nil {
			node = node.left
		} else if node.key < searchKey && node.right != nil {
			node = node.right
		} else {
			return true, node
		}
	}
	return false, node
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

func inOrderHelper[K constraints.Ordered, T any](root *binarySearchTreeNode[K, T], inOrderContainer *([]*binarySearchTreeNode[K, T])) {
	if root == nil {
		return
	}
	inOrderHelper(root.left, inOrderContainer)
	*inOrderContainer = append((*inOrderContainer), root)
	inOrderHelper(root.right, inOrderContainer)
}

func (bst *BinarySearchTree[K, T]) inOrder() []*binarySearchTreeNode[K, T] {
	inOrderContainer := make([]*binarySearchTreeNode[K, T], 0)
	inOrderHelper(bst.root, &inOrderContainer)
	return inOrderContainer
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
