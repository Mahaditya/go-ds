package main

import (
	"fmt"
	"go-ds/src/dstructs"
	"go-ds/src/dstructs/queue"
)

func call(st dstructs.Queue[string]){
	st.Push("lambda")
}

func main() {
	var myQueue queue.Queue[string]
	myQueue.Push("alpha")
	myQueue.Push("beta")
	myQueue.Push("gamma")
	fmt.Println(myQueue)
	
}
