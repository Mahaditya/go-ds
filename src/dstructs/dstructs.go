package dstructs

type Sizeable interface {
	Size() int
	IsEmpty() bool
	IsNotEmpty() bool
}

type Indexable[T any] interface {
	At(int)(T,error)
}

type Popable[T any] interface {
	Pop() (T,error)
}

type Pushable[T any] interface {
	Push(el T)
}

type TopPeekable[T any] interface {
	Top() (T,error)
}

type FrontPeekable[T any] interface {
	Front() (T,error)
}

type EndAccessible[T any] interface {
	First() (T,error)
	Last() (T,error)
}

type Searchable[T any] interface {
	Contains(T) bool
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
	FrontPeekable[T]
	Sizeable
}

type Vector[T any] interface {
	Popable[T]
	Pushable[T]
	Indexable[T]
	EndAccessible[T]
	Searchable[T]
	Sizeable
}

