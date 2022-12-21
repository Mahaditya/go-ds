package dstructs

type Sizeable interface {
	Size() int
	IsEmpty() bool
	IsNotEmpty() bool
}

type Indexable[T any] interface {
	At(int) (T, error)
}

type ListIndexable[T any] interface {
	Get(int) (T,error)
	Set(int,T) error
}

type Popable[T any] interface {
	Pop() (T, error)
}

type Pushable[T any] interface {
	Push(el T)
}

type TopPeekable[T any] interface {
	Top() (T, error)
}

type FrontAccessible[T any] interface {
	Front() (T, error)
}

type BackAccessible[T any] interface {
	Back() (T, error)
}

type Searchable[T any] interface {
	Contains(T) bool
	Find(T) (index int, present bool)
}

type Iterator[T any] interface {
	HasNext() bool
	GetNext() (T, error)
}
type Iterable[T any] interface {
	Sizeable
	GetIterator() Iterator[T]
}

type FrontPopable[T any] interface {
	PopFront() (T, error)
}

type Addable[T any] interface {
	Add(T) bool
	AddAll(Iterable[T]) bool
}

type Clearable[T any] interface {
	Clear()
}

type Comparable[T any] interface {
	Equals(any) bool
}

type Removable[T any] interface {
	Remove(T) bool
	RemoveIf(f func(T) bool) bool 
}

type ArrayConvertible[T any] interface{
	ToArray() []T
}

type IndexAddable[T any] interface {
	AddAt(int,T) error
	AddAllAt(int,Iterable[T]) error
}

type IndexRemoveable[T any] interface {
	RemoveAt(int,T) error
}

type Collection[T any] interface {
	Addable[T]
	Removable[T]
	Comparable[T]
	Clearable[T]
	Iterable[T]
	Searchable[T]
	ArrayConvertible[T]
	Sizeable
}

type List[T any] interface{
	Collection[T]
	ListIndexable[T]
	IndexAddable[T]
	IndexRemoveable[T]
}

type Stack[T any] interface {
	Popable[T]
	Pushable[T]
	Sizeable
	TopPeekable[T]
}

type Queue[T any] interface {
	Popable[T]
	Pushable[T]
	FrontAccessible[T]
	Sizeable
}

type Vector[T any] interface {
	Popable[T]
	Pushable[T]
	Indexable[T]
	FrontAccessible[T]
	BackAccessible[T]
	Searchable[T]
	Iterable[T]
	FrontPopable[T]
	Sizeable
}

type SinglyLinkedList[T any] interface {
	Sizeable
	Pushable[T]
	Indexable[T]
	Popable[T]
}
