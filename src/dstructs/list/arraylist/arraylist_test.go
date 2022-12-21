package arraylist

import (
	"reflect"
	"testing"
)

func TestAddAt(t *testing.T) {
	var al ArrayList[int]
	err := al.AddAt(0, 1)

	if err != errIndexOutOfBounds {
		t.Error("Should have thrown error")
	}

	al.data = []int{0, 1, 2, 3, 4, 5}

	tests := []struct {
		input []int
		want  ArrayList[int]
		err   bool
	}{
		{[]int{0, 1}, *From(1, 0, 1, 2, 3, 4, 5), false},
		{[]int{1, 5}, *From(1, 5, 0, 1, 2, 3, 4, 5), false},
		{[]int{11, 9}, *From(1, 5, 0, 1, 2, 3, 4, 5), true},
	}

	for _, val := range tests {
		err := al.AddAt(val.input[0], val.input[1])
		if err != nil && !val.err {
			t.Error("should have thrown error")
			continue
		}
		got := al
		if !reflect.DeepEqual(got, val.want) {
			t.Errorf("got %v, wanted %v", got, val.want)
		} else {
			t.Log("Passed")
		}
	}

}

