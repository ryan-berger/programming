package main

import (
	"fmt"
	"github.com/ryan-berger/programming/a-and-d/sorts"
)

var (
	data = []int {4, 1, 3, 5, 2, 0 , -1}
	copyOfData = []int {4, 1, 3, 5, 2, 0 , -1}
)

func reset()  {
	copy(data, copyOfData)
}

func main()  {
	sorts.InsertionSort(data)
	fmt.Println("Sorted", data)
	reset()
	fmt.Println("Unsort", data)
}