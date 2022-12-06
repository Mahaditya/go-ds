package vector

import (
	"testing"
)

func TestSize(t *testing.T){
	v:=new(Vector[int])
	
    got := v.Size()
    want := 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

    v.Push(1).Push(2)

    got= v.Size()
    want= 2
    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

    v.Pop()

    got=v.Size()
    want=1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

    v.Pop()

    got=v.Size()
    want=0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

    _,err1:=v.Pop()

    if err1==nil{
        t.Error("should have thrown error")
    }

    _,err2:=v.Pop()

    if err2==nil{
        t.Error("should have thrown error")
    }

}

func TestAt(t *testing.T){
    v:=new(Vector[int])

    _,err:=v.At(1)
    if err==nil{
        t.Error("should have thrown error")
    }



    v.Push(1).Push(2)

    got,_:=v.At(0)
    want:= 1

    if got!=want{
        t.Errorf("got %q, wanted %q", got, want)
    }

    _,err=v.At(-1)

    if err==nil{
        t.Error("should have thrown error")
    }

}


func TestPopFront(t *testing.T){
    v:=new(Vector[int])
    v.Push(1).Push(2)
    got,_:=v.PopFront()
    want:= 1
    if got!=want{
        t.Errorf("got %q, wanted %q", got, want)
    }
    got,_=v.PopFront()
    want=2
    if got!=want{
        t.Errorf("got %q, wanted %q", got, want)
    }
    _,err:=v.PopFront()

    if err==nil{
        t.Error("should have thrown error")
    }
}