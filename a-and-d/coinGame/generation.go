package coinGame

import (
	"fmt"
	"github.com/ryan-berger/programming/a-and-d/graphs"
)

// node cache for when we build out the graph
var nodeMap = map[CoinString]*graphs.GraphNode{

}

// initial coin list
var initialString = CoinList {
	H,
	T,
	T,
	H,
}

// get a node or create it, returning whether either the new
// node or cached node. If it is new we return a bool saying it
// is so that the build tree function knows not to go searching for the
// node's children
func getOrCreate(list CoinList) (node *graphs.GraphNode, new bool)  {
	key := list.String()
	if value, ok := nodeMap[key]; ok {
		return value, !ok
	} else {
		node := graphs.NewNode(list)
		nodeMap[key] = node
		return node, !ok
	}
}

// build a graph with the root node.
func buildGraph(node *graphs.GraphNode)  {
	// cast the value since the graph's implementation is
	// "generic"
	value := node.Value.(CoinList)
	permutations := value.GetPermutations()

	for _, permutation := range permutations {
		// get node and check to see if node already has been searched
		cachedNode, exists := getOrCreate(permutation)

		// is the move made reversible?
		if value.IsReversible(permutation) {
			// if yes, then we want parent <-> child
			node.AddTwoWayEdge(1, node, cachedNode)
		} else {
			// otherwise we only want a parent -> child
			node.AddOneWayEdge(1, cachedNode)
		}

		// if the node already is in the map, don't search it
		// because if it exists in the map, it has been, or is
		// about to be searched
		if exists {
			buildGraph(cachedNode)
		}
	}
}

func GenerateNodes() {
	rootNode := graphs.NewNode(initialString)
	buildGraph(rootNode)
	i := 0
	for range nodeMap {
		i++
	}
	fmt.Println("num nodes:", i)
}