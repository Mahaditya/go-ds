package dstructs_test

import (
	"fmt"
	"go-ds/src/dstructs"
	"go-ds/src/dstructs/list/arraylist"
	"go-ds/src/dstructs/sll"
	"go-ds/src/dstructs/vector"
	"testing"
)

func TestVectorInterface(t *testing.T) {
	var vect dstructs.Vector[int]
	vectStruct := vector.From(1, 3, 5)
	vect = &vectStruct
	fmt.Println(vect)
}

func TestSllInterface(t *testing.T){
	var Isll dstructs.SinglyLinkedList[int]
	slStruct:= sll.Make(1)
	Isll = slStruct
	fmt.Println(Isll)
}

func TestListInterface(t *testing.T){
	var arrayList dstructs.List[int]
	al:= arraylist.From(1,3,5)
	arrayList = al
	fmt.Println(arrayList)
}