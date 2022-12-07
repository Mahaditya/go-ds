package hashset

type HashSet[T comparable] struct {
	data map[T]bool
}

func (hs *HashSet[T]) Insert(element T) *HashSet[T] {
	if hs.data == nil {
		hs.data = make(map[T]bool)
	}
	hs.data[element] = true
	return hs
}

func (hs *HashSet[T]) Contains(element T) bool {
	_, ok := hs.data[element]
	return ok
}

func (hs *HashSet[T]) ForEach(f func(key T)) {
	for key := range hs.data {
		f(key)
	}
}

func (hs *HashSet[T]) Remove(key T){
	delete(hs.data,key)
}