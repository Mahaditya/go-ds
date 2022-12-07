package vector

import (
	"errors"
	"fmt"
)

type Vector[T any] struct {
	data []T
}

func Make[T any](slice []T) Vector[T]{
	v:= new(Vector[T])
	v.data = slice
	return *v
}

func (v *Vector[T]) Size() int {
	return len(v.data)
}

func (v *Vector[T]) Push(element T) *Vector[T] {
	v.data = append(v.data, element)
	return v
}

func (v *Vector[T]) IsEmpty() bool {
	return v.Size() <= 0
}

func (v *Vector[T]) IsNotEmpty() bool {
	return !v.IsEmpty()
}

func (v *Vector[T]) Pop() (T, error) {
	val, err := v.Last()
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

func (v *Vector[T]) Last() (T, error) {
	return v.At(v.Size() - 1)
}

func (v *Vector[T]) First() (T, error) {
	return v.At(0)
}

func (v *Vector[T]) PopFront() (T, error) {
	val, err := v.First()
	if err == nil {
		v.data = v.data[1:]
	}
	return val, err
}

func (v *Vector[T]) Print() {
	fmt.Println(v.data)
}
