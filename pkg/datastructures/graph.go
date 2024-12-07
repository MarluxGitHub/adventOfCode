package datastructures

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Name  string
	Edges []*Edge
	Value string
}

type Edge struct {
	To   *Node
	Cost int
}

func (g *Graph) GetNode(name string) *Node {
	for _, n := range g.Nodes {
		if n.Name == name {
			return n
		}
	}

	return nil
}

func (g *Graph) AddNode(name string, value string) *Node {
	n := &Node{
		Name:  name,
		Edges: make([]*Edge, 0),
		Value: value,
	}

	g.Nodes = append(g.Nodes, n)
	return n
}

func (g *Graph) GetNodes() []*Node {
	return g.Nodes
}

func (n *Node) AddEdge(to *Node, cost int) {
	n.Edges = append(n.Edges, &Edge{To: to, Cost: cost})
}

func (n *Node) GetEdges() []*Edge {
	return n.Edges
}

func (n *Node) GetEdge(to *Node) *Edge {
	for _, e := range n.Edges {
		if e.To == to {
			return e
		}
	}

	return nil
}

func NewGraph() *Graph {
	return &Graph{Nodes: make([]*Node, 0)}
}
