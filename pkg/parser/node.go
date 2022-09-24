package parser

type NodeKind interface {
	String() string
}

type Node struct {
	Kind     NodeKind
	LHS      *Node
	RHS      *Node
	Child    *Node
	Children []*Node

	S string
	F float64
	I int64
	B bool
}

func (n *Node) Clone() *Node {
	return NewNode(n.Kind, n.LHS, n.RHS, n.Child, n.Children, n.S, n.F, n.I, n.B)
}

func NewNode(kind NodeKind, lhs, rhs, child *Node, children []*Node, s string, f float64, i int64, b bool) *Node {
	return &Node{
		Kind:     kind,
		LHS:      lhs,
		RHS:      rhs,
		Child:    child,
		Children: children,
		S:        s,
		F:        f,
		I:        i,
		B:        b,
	}
}
