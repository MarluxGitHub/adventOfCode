package datastructures

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Name  string
	Edges []*Edge
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

func (g *Graph) AddNode(name string) *Node {
	n := &Node{Name: name}
	g.Nodes = append(g.Nodes, n)
	return n
}

func (g *Graph) AddEdge(from, to *Node, cost int) {
	e := &Edge{To: to, Cost: cost}
	from.Edges = append(from.Edges, e)
}

func NewGraph() *Graph {
	return &Graph{Nodes: make([]*Node, 0)}
}
