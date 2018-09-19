package main

import "github.com/ryan-berger/programming/a-and-d/coinGame"

func main()  {
	root, end := coinGame.GenerateNodes()
	coinGame.FindShortestPath(root, end)
}
