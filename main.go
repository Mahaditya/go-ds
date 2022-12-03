package main

import (
	"fmt"
	"go-ds/src/dstructs"
)

func main(){
	bst:= new(dstructs.BinarySearchTree[int,string])
	bst.Insert(1,"a").Insert(5,"b").Insert(2,"c")
	fmt.Println(bst.Search(2))

}