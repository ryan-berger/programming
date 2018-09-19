package coinGame

import (
	"github.com/ryan-berger/programming/a-and-d/graphs"
	"fmt"
)

var alreadySearched = map[CoinString]*graphs.GraphNode {

}

type pathNode struct {
	node *graphs.GraphNode
	previous *pathNode
}

func FindShortestPath(start, last *graphs.GraphNode)  {
	fmt.Println("start:", start.Value.(CoinList).String())
	fmt.Println("end:", last.Value.(CoinList).String())

	fmt.Printf("start: %p \n", start.Value.(CoinList))
	fmt.Printf("end: %p \n", last.Value.(CoinList))
	startSearch(start, last)
}

func nodesFromEdges(previous *pathNode, edges []*graphs.Edge) (nodes []*pathNode) {
	for _, edge := range edges {
		nodes = append(nodes, &pathNode{node: edge.Node, previous: previous})
	}
	return nodes
}


func hasBeenSearched(node *graphs.GraphNode) bool  {
	_, ok := alreadySearched[node.Value.(CoinList).String()]
	return ok
}

func addToSearch(node *graphs.GraphNode)  {
	alreadySearched[node.Value.(CoinList).String()] = node
}

func areSame(node *graphs.GraphNode, path *pathNode) bool  {
	return node.Value.(CoinList).IsSame(path.node.Value.(CoinList))
}

func startSearch(start, end *graphs.GraphNode)  {
	found := false
	currentNode := start
	children := nodesFromEdges(nil, currentNode.Edges)


	i := 0
	for !found && i < 20 {
		fmt.Println("depth: ", i)
		var nextLevel [][]*pathNode
		for _, child := range children {
			if !hasBeenSearched(child.node) && areSame(end, child) {
				nextLevel = append(nextLevel, nodesFromEdges(child, child.node.Edges))
				addToSearch(child.node)
			} else if areSame(end, child) {
				fmt.Println("yeet")
				found = true
				getPath(child)
				goto FINISH
			}
		}
		children = flatten(nextLevel)
		i++
	}
	FINISH:
}

func getPath(node *pathNode)  {
	path := []*graphs.GraphNode{node.previous.node}

	currentNode := node

	for currentNode.previous != nil {
		path = append(path, currentNode.previous.node)
		currentNode = currentNode.previous
	}

	for _, p := range path {
		fmt.Println(p.Value.(CoinList).String())
	}
}

func flatten(list [][]*pathNode) (flattened []*pathNode)  {
	for _, item := range list {
		for _, inner := range item {
			flattened = append(flattened, inner)
		}
	}
	return flattened
}