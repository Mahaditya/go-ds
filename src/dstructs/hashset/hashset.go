package hashset

import (
	"fmt"
	"go-ds/src/dstructs/hashmap"
)

type HashSet[T comparable] struct {
	mp hashmap.HashMap[T, bool]
}

func (hs *HashSet[T]) Insert(element T) *HashSet[T] {
	hs.mp.Insert(element, true)
	return hs
}

func (hs *HashSet[T]) Contains(element T) bool {
	return hs.mp.ContainsKey(element)
}

func (hs *HashSet[T]) ForEach(f func(key T)) {
	hs.mp.ForEach(func(key T, value bool) { f(key) })
}

func (hs *HashSet[T]) Remove(key T) {
	hs.mp.Remove(key)
}

func (hs *HashSet[T]) Size() int {
	return hs.mp.Size()
}

func (hs HashSet[T]) Format(f fmt.State, verb rune) {
	fmt.Print("{ ")
	hs.mp.ForEach(func(key T, value bool) {
		fmt.Print(key, " ")
	})
	fmt.Print("}")
}
