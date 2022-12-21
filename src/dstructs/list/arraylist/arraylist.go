package arraylist

import (
	"errors"
	"go-ds/src/dstructs"
	"reflect"
)

type ArrayList[T any] struct {
	data []T
}

var errIndexOutOfBounds = errors.New("index out of bounds")

func From[T any](elements ...T) *ArrayList[T] {
	var arrayList ArrayList[T]
	arrayList.data = elements
	return &arrayList
}

func FromIterable[T any](iterable dstructs.Iterable[T]) *ArrayList[T]{
	var arrayList ArrayList[T]
	it:= iterable.GetIterator()
	for it.HasNext(){
		if next,err:=it.GetNext();err==nil{
			arrayList.Add(next)
		}
	}
	return &arrayList
}

func FromArray[T any](array []T) *ArrayList[T]{
	var arrayList ArrayList[T]
	copy(arrayList.data,array)
	return &arrayList
}

func (al *ArrayList[T]) Add(element T) bool {
	al.data = append(al.data, element)
	return true
}

func (al *ArrayList[T]) AddAt(index int,element T) error {
	if !al.isValidIndex(index){
		return errIndexOutOfBounds
	}	
	al.data =append(al.data[:index],append(From(element).data,al.data[index:]...)...)
	return nil
}

func (al *ArrayList[T]) AddAll(iterable dstructs.Iterable[T]) bool {
	arrayList:= FromIterable(iterable)
	al.data = append(al.data, arrayList.data...)
	return true
}


func (al *ArrayList[T]) AddAllAt(index int,iterable dstructs.Iterable[T]) error {
	if !al.isValidIndex(index){
		return errIndexOutOfBounds
	}
	arrayList:=FromIterable(iterable)
	al.data = append(al.data[:index],append(arrayList.data,al.data[index:]...)...)
	return nil
}

func (al *ArrayList[T]) Size() int {
	return len(al.data)
}

func (al *ArrayList[T]) IsEmpty() bool {
	return al.Size()<=0;
}

func (al *ArrayList[T]) IsNotEmpty() bool {
	return !al.IsEmpty()
}

func (al *ArrayList[T]) GetIterator() dstructs.Iterator[T] {
	return &ArrayListIterator[T]{0,al}
}

func (al *ArrayList[T]) Get(index int) (T,error) {
	if !al.isValidIndex(index) {
		return *new(T), errIndexOutOfBounds
	}
	return al.data[index], nil
}

func (al *ArrayList[T]) Set(index int,element T) (error) {
	if !al.isValidIndex(index) {
		return errIndexOutOfBounds
	}
	al.data[index] = element
	return nil
}

func (al *ArrayList[T]) Clear(){
	al.data = make([]T, 0)
}

func (al *ArrayList[T]) Find(element T) (index int, present bool) {
	for idx, val := range al.data {
		if reflect.DeepEqual(element, val) {
			return idx, true
		}
	}
	return -1, false
}


func (al *ArrayList[T]) Contains(element T) bool {
	_,present:= al.Find(element)
	return present
}

func (al *ArrayList[T]) Equals(object any) bool {
	return reflect.DeepEqual(al,object)
}

func (al *ArrayList[T]) Remove(element T) bool {
	if al.Size()<=0 {
		return false
	}
	al.data = al.data[:al.Size()-1]
	return true
}

func (al *ArrayList[T]) RemoveIf(f func(T)bool) bool {
	var arr []T
	for _,val:=range al.data{
		if !f(val){
			arr = append(arr, val)
		}
	}
	if al.Size() > len(arr){
		al.data = arr
		return true
	}
	return false
}

func (al *ArrayList[T]) ToArray() []T{
	return al.data
}