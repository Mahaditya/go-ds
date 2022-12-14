package bst

import (
	"go-ds/src/dstructs/queue"

	"golang.org/x/exp/constraints"
)

type bstNode[K constraints.Ordered, T any] struct {
	left  *bstNode[K, T]
	right *bstNode[K, T]
	key   K
	value T
}

func (bst *BinarySearchTree[K, T]) search(searchKey K) (found bool, value *bstNode[K, T]) {
	if bst.root == nil {
		return false, bst.root
	}
	node := bst.root
	for node != nil {
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

func (bst *BinarySearchTree[K, T]) levelOrderTraversal() [][]*bstNode[K, T] {
	levelOrder := make([][]*bstNode[K, T], 0)
	q := new(queue.Queue[*bstNode[K, T]])
	q.Push(bst.root)
	for q.IsNotEmpty() {
		thisLevel := make([]*bstNode[K, T], 0)
		for n := q.Size(); n > 0; n-- {
			node, _ := q.Pop()
			thisLevel = append(thisLevel, node)
			if node.left != nil {
				q.Push(node.left)
			}
			if node.right != nil {
				q.Push(node.right)
			}
		}
		levelOrder = append(levelOrder, thisLevel)
	}
	return levelOrder
}

func inOrderHelper[K constraints.Ordered, T any](root *bstNode[K, T], inOrderContainer *([]*bstNode[K, T])) {
	if root == nil {
		return
	}
	inOrderHelper(root.left, inOrderContainer)
	*inOrderContainer = append((*inOrderContainer), root)
	inOrderHelper(root.right, inOrderContainer)
}

func (bst *BinarySearchTree[K, T]) inOrder() []*bstNode[K, T] {
	inOrderContainer := make([]*bstNode[K, T], 0)
	inOrderHelper(bst.root, &inOrderContainer)
	return inOrderContainer
}
