package pair

type Pair[A, B any] struct {
	first  A
	second B
}

func Make[A, B any](first A, second B) Pair[A, B] {
	pair := new(Pair[A, B])
	pair.first = first
	pair.second = second
	return *pair
}

func (p Pair[A, B]) Get() (A, B) {
	return p.first, p.second
}
