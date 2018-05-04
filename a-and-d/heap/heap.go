package heap

import "fmt"

type Node int
type Location int
type Heap []Node

type Children struct {
	left Location
	right Location
}


func (heap *Heap) Insert(node Node)  {
	*heap = append(*heap, node)
	heap.Heapify()
}

func (heap *Heap) GetParent(index Location) Location {
	return Location((index - 1) / 2)
}

func (heap *Heap) GetChildren(index Location) *Children  {
	return &Children{
		Location(index * 2),
		Location((index * 2) + 1),
	}

}

func (heap *Heap) Get(index Location) Node {
	return (*heap)[int(index)]
}

func (heap *Heap) Heapify() {
	h := *heap
	currentIndex := Location(len(h) - 1)
	fmt.Println()
	for h[currentIndex] > h[h.GetParent(currentIndex)]  && currentIndex > 0 {
		h.swap(currentIndex, h.GetParent(currentIndex))
		currentIndex = h.GetParent(currentIndex)
	}
}

func (heap *Heap) swap(a, b Location) {
	temp := (*heap)[a]
	(*heap)[a] = (*heap)[b]
	(*heap)[b] = temp
}
