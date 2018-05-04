package graphs

type Edge struct {
	Parent *GraphNode
	Node   *GraphNode
	Cost   int
}

func NewOneWayEdge(cost int, node *GraphNode) *Edge  {
	return &Edge{
		Parent: nil,
		Node: node,
		Cost: cost,
	}
}

func NewTwoWayEdge(cost int, parent, child *GraphNode) *Edge  {
	return &Edge{
		Parent: parent,
		Node: child,
		Cost: cost,
	}
}

type GraphNode struct {
	Value interface{}
	Edges []*Edge
}

func (node *GraphNode) AddOneWayEdge(cost int, edgeNode *GraphNode) {
	node.Edges = append(node.Edges, NewOneWayEdge(cost, edgeNode))
}

func (node *GraphNode) AddTwoWayEdge(cost int, parent, child *GraphNode)  {
	node.Edges = append(node.Edges, NewTwoWayEdge(cost, parent, child))
}

type Graph struct {
	Nodes []*GraphNode
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: []*GraphNode{},
	}
}

func NewNode(value interface{}) *GraphNode {
	return &GraphNode{
		Value: value,
	}
}
