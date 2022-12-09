package hashmap

import "fmt"

type HashMap[T comparable, U any] struct {
	data map[T]U
}

func (hs *HashMap[T, U]) Insert(key T, value U) *HashMap[T, U] {
	if hs.data == nil {
		hs.data = make(map[T]U)
	}
	hs.data[key] = value
	return hs
}

func (hs *HashMap[T, U]) ContainsKey(key T) bool {
	_, ok := hs.data[key]
	return ok
}

func (hs *HashMap[T, U]) ForEach(f func(key T, value U)) {
	for key, value := range hs.data {
		f(key, value)
	}
}

func (hs *HashMap[T, U]) Remove(key T) {
	delete(hs.data, key)
}

func (hs *HashMap[T, U]) Size() int {
	return len(hs.data)
}

func (hs *HashMap[T, U]) Get(key T) (U, bool) {
	val, ok := hs.data[key]
	return val, ok
}

func (hs *HashMap[T, U]) Replace(key T, value U) *HashMap[T, U] {
	containsKey := hs.ContainsKey(key)
	if containsKey {
		hs.data[key] = value
	}
	return hs
}

func (hs *HashMap[T, U]) Update(key T,f func(U)U) *HashMap[T, U] {
	hs.Insert(key,f(hs.data[key]))
	return hs
}

func (hs HashMap[T, U]) String() string {
	return fmt.Sprintf("%v",hs.data)
}
