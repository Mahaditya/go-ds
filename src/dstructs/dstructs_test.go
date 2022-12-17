package dstructs_test

import (
	"fmt"
	"go-ds/src/dstructs"
	"go-ds/src/dstructs/vector"
	"testing"
)

func TestInterface(t *testing.T) {
	var vect dstructs.Vector[int]
	vectStruct := vector.From(1, 3, 5)
	vect = &vectStruct
	fmt.Println(vect)
}
