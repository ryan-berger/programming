package main

import (
	"fmt"
	"sort"
)

type stops []int

func (s stops) Len() int {
	return len(s)
}
func (s stops) Less(i, j int) bool  {
	return s[i] > s[j]
}

func (s stops) Swap(i, j int)  {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

type stop struct {
	stops []int
}

func (s stop) String() string {
	if len(s.stops) != 1 {
		return fmt.Sprintf("%d-%d", s.stops[0], s.stops[len(s.stops) - 1])
	}
	return fmt.Sprintf("%d", s.stops[0])
}

func main()  {
	var inputNumber int
	fmt.Scanln(&inputNumber)

	var s stops

	for i := 0; i < inputNumber; i++ {
		var stopNumber int
		fmt.Scan(&stopNumber)
		s = append(s, stopNumber)
	}

	sort.Sort(s)

	lastNumber := s[0]
	var formattedStops []*stop

	for i := 1; i < len(s); i++ {
		var stops []int
		if s[i] - 1 == lastNumber {
			for j := 0; j < len(s); j++ {
				
				i++
			}
		}
		newStop := &stop{
			stops: stops,
		}
		formattedStops = append(formattedStops, newStop)
	}
}
