package vector

import (
	"errors"
	"fmt"
)

type Vector[T any] struct {
	data []T
}

func From[T any](slice []T) Vector[T]{
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

func Map[A,B any](vector Vector[A],f func (A)B) Vector[B] {
	newVector:= new(Vector[B])
	for _,val:=range vector.data{
		newVector.Push(f(val))
	}
	return *newVector
}

func (v *Vector[T]) Replace(index int, element T) (*Vector[T],error){
	_,err:=v.At(index)
	if err==nil{
		v.data[index] = element
	}
	return v,err
}

func (v Vector[T]) SubVector(startIndex int, endIndex int) (Vector[T],error){
	_,err:=v.At(startIndex)
	_,err2:=v.At(endIndex)

	if err==nil && err2==nil{
		return From(v.data[startIndex:endIndex]),nil
	}

	if err!=nil{
		return *new(Vector[T]),err
	}
	
	return *new(Vector[T]),err2

}

func (v Vector[T]) String() string{
	return fmt.Sprintf("<Vector %v>",v.data)
}

func (v Vector[T]) ToString() string {
	return fmt.Sprintf("%v",v.data)
}