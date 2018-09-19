package main

import (
	"fmt"
	"sync"
	"math"
)

type point struct {
	x int64
	y int64
}

type workRequest struct {
	p      *point
	points *[]*point
}

var count = 0
var countMu = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func main() {
	var input int
	fmt.Scanln(&input)

	points := &[]*point{}
	wg.Add(input)
	wr := make(chan workRequest, input)
	startWorkers(wr)

	for i := 0; i < input; i ++ {
		p := &point{}
		fmt.Scanf("%d %d", &p.x, &p.y)
		wr <- workRequest{p: p, points: points}
		*points = append(*points, p)
	}

	wg.Wait()
	fmt.Println(count)
}

func startWorkers(wr chan workRequest) {
	for i := 0; i < 300; i++ {
		go readInWork(wr)
	}
}

func readInWork(wr chan workRequest) {
	for {
		select {
		case request := <-wr:
			run(request.p, request.points)
		}
	}
}

func run(p *point, points *[]*point) {
	deRef := *points
	fmt.Println(deRef)
	for j := 0; j < len(deRef); j++ {
		xDiff := int(math.Abs(float64(p.x) - float64((deRef)[j].x)))
		yDiff := int(math.Abs(float64(p.y) - float64((deRef)[j].y)))

		if (xDiff == 0 && yDiff == 2018) ||
			(xDiff == 2018 && yDiff == 0) ||
			(xDiff == 1118 && yDiff == 1680) ||
			(xDiff == 1680 && yDiff == 1118) {
			countMu.Lock()
			count++
			countMu.Unlock()
		}
	}
	wg.Done()
}
