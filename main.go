package main

import (
	"fmt"
	"go-ds/src/dstructs/bst"
	"math/rand"
)

func main(){
	bst:= new(bst.BinarySearchTree[int,int])
	for i:=0;i<1000;i++{
		bst.Insert(rand.Intn(1000),rand.Intn(1000))
	}
	fmt.Println(bst.Search(2))
	bst.Print()
}