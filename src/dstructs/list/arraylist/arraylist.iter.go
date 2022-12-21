package arraylist

import "errors"

type ArrayListIterator[T any] struct {
	index  int
	arrayList *ArrayList[T]
}

func (ali *ArrayListIterator[T]) HasNext() bool {
	return ali.index < ali.arrayList.Size()
}

func (ali *ArrayListIterator[T]) GetNext() (T, error) {
	if ali.HasNext() {
		if next, err := ali.arrayList.Get(ali.index); err != nil {
			return next, err
		} else {
			ali.index++
			return next, err
		}
	}
	return *new(T), errors.New("no next element")
}
