package vector

import (
	"errors"
	"fmt"
	"go-ds/src/dstructs"
	"reflect"
)

type Vector[T any] struct {
	data []T
}

func From[T any](slice []T) Vector[T] {
	v := new(Vector[T])
	v.data = slice
	return *v
}

func (v *Vector[T]) Size() int {
	return len(v.data)
}

func (v *Vector[T]) Push(element T) {
	v.data = append(v.data, element)
}

func (v *Vector[T]) IsEmpty() bool {
	return v.Size() <= 0
}

func (v *Vector[T]) IsNotEmpty() bool {
	return !v.IsEmpty()
}

func (v *Vector[T]) Pop() (T, error) {
	val, err := v.Back()
	if err == nil {
		v.data = v.data[:v.Size()-1]
	}
	return val, err
}

func (v *Vector[T]) At(index int) (T, error) {
	if index < 0 || index >= v.Size() {
		return *new(T), errors.New("index out of bounds")
	}
	return v.data[index], nil
}

func (v *Vector[T]) Back() (T, error) {
	return v.At(v.Size() - 1)
}

func (v *Vector[T]) Front() (T, error) {
	return v.At(0)
}

func (v *Vector[T]) PopFront() (T, error) {
	val, err := v.Front()
	if err == nil {
		v.data = v.data[1:]
	}
	return val, err
}

func Map[A, B any](vector Vector[A], f func(A) B) Vector[B] {
	newVector := new(Vector[B])
	for _, val := range vector.data {
		newVector.Push(f(val))
	}
	return *newVector
}

func (v Vector[T]) SubVector(startIndex int, endIndex int) (Vector[T], error) {
	_, err := v.At(startIndex)
	_, err2 := v.At(endIndex)

	if err == nil && err2 == nil {
		return From(v.data[startIndex:endIndex]), nil
	}

	if err != nil {
		return *new(Vector[T]), err
	}

	return *new(Vector[T]), err2

}

func (v Vector[T]) String() string {
	return fmt.Sprintf("<Vector %v>", v.data)
}

func (v Vector[T]) ToString() string {
	return fmt.Sprintf("%v", v.data)
}

func (v *Vector[T]) GetIterator() dstructs.Iterator[T] {
	return &vectorIterator[T]{0, v}
}

func (v *Vector[T]) Contains(element T) bool {
	_, present := v.Find(element)
	return present
}

func (v *Vector[T]) Find(element T) (index int, present bool) {
	for idx, val := range v.data {
		if reflect.DeepEqual(element, val) {
			return idx, true
		}
	}
	return -1, false
}

func (v *Vector[T]) ForEach(f func(T)){
	for _,val:= range v.data{
		f(val)
	}
}

func Make[T any](size int) Vector[T]{
	slc := make([]T,size)
	return From(slc)
}

func (v *Vector[T]) Replace(idx int,val T) bool {
	if _,err:= v.At(idx); err!=nil{
		return false
	}
	v.data[idx] = val
	return true
}