package prefix_tree

import "SignatureScanner/internal/signature"

type Node struct {
	Point      string
	nodes      []*Node
	Signatures map[string]*signature.Signature
}

func NewNode(point string, signat string) *Node {
	nd := &Node{
		Point:      point,
		Signatures: make(map[string]*signature.Signature),
	}
	if signat != "" {
		nd.Signatures[signat] = signature.NewSignature(signat)
	}
	return nd
}

func (nd *Node) ContainsSi(siStr string) bool {
	_, ok := nd.Signatures[siStr]
	return ok
}
