package coinGame
//
//import (
//	"github.com/ryan-berger/programming/a-and-d/graphs"
//	"fmt"
//)
//
//var alreadySearched = map[CoinString]*graphs.GraphNode {
//
//}
//
//func FindPath(start, last *graphs.GraphNode)  {
//	fmt.Println("start:", start.Value.(CoinList).String())
//	fmt.Println("end:", last.Value.(CoinList).String())
//	path := search(start, last, []*graphs.GraphNode{})
//	fmt.Print(path)
//}
//
//func hasBeenSearched(node *graphs.GraphNode) bool  {
//	_, ok := alreadySearched[node.Value.(CoinList).String()]
//	return ok
//}
//
//func addToSearch(node *graphs.GraphNode)  {
//	alreadySearched[node.Value.(CoinList).String()] = node
//}
//
//
//func search(current, end *graphs.GraphNode, path []*graphs.GraphNode) []*graphs.GraphNode {
//	newPath := make([]*graphs.GraphNode, len(path))
//	copy(newPath, path)
//
//	newPath = append(newPath, current)
//
//	fmt.Println("depth:", len(newPath))
//
//	searchNodes := nodesFromEdges(current.Edges)
//	for _, node := range searchNodes {
//		fmt.Println("node:", node.Value.(CoinList).String(), "has been searched:", hasBeenSearched(node))
//		if !hasBeenSearched(node) {
//			if node == end {
//				return path
//			}
//		}
//	}
//
//	for _, node := range searchNodes {
//		search(node, end, newPath)
//		addToSearch(node)
//	}
//	return nil
//}
//
//
//
//func nodesFromEdges(edges []*graphs.Edge) (nodes []*graphs.GraphNode) {
//	for _, edge := range edges {
//		nodes = append(nodes, edge.Node)
//	}
//	return nodes
//}
//
