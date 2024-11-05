package tree

// Node is a struct containing nodes of a file tree.
type Node struct {
	// can be nil
	Parent *Node
	Name   string
	// Empty value mean directory
	Ext string
	// cannot be nil; may have len 0
	Children []*Node
}

// NewNode constructs and returns a Node instance.
// parent and children are optional and can be nil,
// in which case Parent will be stored as nil,
// but Children will be initialized as []*Node of len 0.
// If ext == "", the Node is a directory.
func NewNode(name string, ext string, parent *Node, children []*Node) *Node {
	var c []*Node
	if children == nil {
		c = make([]*Node, 0)
	} else {
		c = children
	}
	n := &Node{
		Name:     name,
		Ext:      ext,
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

func (n *Node) IsDir() bool {
	return n.Ext == ""
}

// TODO: Implement recursive depth-first traversal to process a tree

func Traverse(root *Node) {
}
