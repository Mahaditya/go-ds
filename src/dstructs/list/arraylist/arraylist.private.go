package arraylist

func (al *ArrayList[T]) isValidIndex(index int)bool{
	if index < 0 || index >= al.Size() {
		return false
	}
	return true
}