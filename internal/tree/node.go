package tree

// Node is a struct containing nodes of a file tree.
type Node struct {
	// can be nil
	Parent   *Node
	Name     string
	FileType string
	// cannot be nil; may have len 0
	Children []*Node
}

// NewNode constructs and returns a Node instance.
// parent and children are optional and can be nil,
// in which case Parent will be stored as nil,
// but Children will be initialized as []*Node of len 0.
func NewNode(name string, ft string, parent *Node, children []*Node) *Node {
	var c []*Node
	if children == nil {
		c = make([]*Node, 0)
	} else {
		c = children
	}
	n := &Node{
		Name:     name,
		FileType: ft,
		Parent:   parent,
		Children: c,
	}
	return n
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node) IsTail() bool {
	return len(n.Children) == 0
}
