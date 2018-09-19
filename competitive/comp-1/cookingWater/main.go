package main

import (
	"fmt"
	"sort"
)

type datapoint struct {
	start int
	end int
}

func (d datapoint) diff() int {
	return d.end - d.start
}

func (d datapoint) inRange(point datapoint) bool  {
	return point.start >= d.start && point.start <= d.end
}

type datapoints []datapoint

func (points datapoints) Len() int {
	return len(points)
}
func (points datapoints) Less(i, j int) bool  {
	return points[i].diff() > points[j].diff()
}

func (points datapoints) Swap(i, j int)  {
	temp := points[i]
	points[i] = points[j]
	points[j] = temp
}


func main()  {
	var numCases int
	fmt.Scanln(&numCases)

	var datapoints datapoints

	for i := 0; i < numCases; i++ {
		d := datapoint{}
		fmt.Scanf("%v %v", &d.start, &d.end)
		datapoints = append(datapoints, d)
	}

	sort.Sort(datapoints)
	currentRange := datapoints[0]

	for i := 1; i < len(datapoints); i++ {
		if currentRange.inRange(datapoints[i]) {
			currentRange = datapoints[i]
		} else {
			fmt.Println("edward is right")
			return
		}
	}

	fmt.Println("gunilla has a point")
}