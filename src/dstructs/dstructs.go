package dstructs

type Sizeable interface {
	Size() int
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