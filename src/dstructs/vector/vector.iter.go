package vector

import "errors"

type vectorIterator[T any] struct {
	index int
	vector *Vector[T]
}

func (vi *vectorIterator[T]) HasNext() bool {
	return vi.index < vi.vector.Size()
}

func (vi *vectorIterator[T]) GetNext() (T,error) {
	if vi.HasNext(){
		if next,err:=vi.vector.At(vi.index);err!=nil{
			return next,err
		}else{
			vi.index++
			return next,err
		}
	}
	return *new(T),errors.New("no next element")
}

