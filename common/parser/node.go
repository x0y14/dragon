package parser

type NodeKind interface {
	String() string
}

type Node struct {
	Kind     NodeKind
	LHS      *Node
	RHS      *Node
	Children []*Node
	N1       *Node
	N2       *Node
	N3       *Node

	S string
	F float64
	I int64
	B bool
}

func (n *Node) Clone() *Node {
	return NewNode(n.Kind, n.LHS, n.RHS, n.Children, n.N1, n.N2, n.N3, n.S, n.F, n.I, n.B)
}

func NewNode(kind NodeKind, lhs, rhs *Node, children []*Node, n1, n2, n3 *Node, s string, f float64, i int64, b bool) *Node {
	return &Node{
		Kind:     kind,
		LHS:      lhs,
		RHS:      rhs,
		Children: children,
		N1:       n1,
		N2:       n2,
		N3:       n3,
		S:        s,
		F:        f,
		I:        i,
		B:        b,
	}
}
