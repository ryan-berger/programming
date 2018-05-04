package main

import (
	"fmt"
	"github.com/ryan-berger/programming/a-and-d/heap"
)

func main()  {
	h := heap.Heap{}
	for i := 0; i < 10; i++ {
		h.Insert(heap.Node(i))
	}
	fmt.Println(h)
}