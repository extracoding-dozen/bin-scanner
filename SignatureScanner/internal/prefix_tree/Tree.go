package prefix_tree

import (
	"SignatureScanner/internal/signature"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var hexTable [256]string

func init() {
	for i := 0; i < 256; i++ {
		hexTable[i] = fmt.Sprintf("%02x", i)
	}
}

type Tree struct {
	nodes []*Node
}

func NewTree() *Tree {
	return &Tree{nodes: nil}
}

func (tr *Tree) Paste(tSi string) {
	raw := signature.Unpack(tSi)
	if len(raw) == 0 {
		return
	}

	var currentNode *Node
	for _, node := range tr.nodes {
		if node.Point == raw[0] {
			currentNode = node
			break
		}
	}

	if currentNode == nil {
		currentNode = NewNode(raw[0], "")
		tr.nodes = append(tr.nodes, currentNode)
	}

	// 2. Спуск по дереву
	for i := 1; i < len(raw); i++ {
		point := raw[i]
		var foundChild *Node

		for _, child := range currentNode.nodes {
			if child.Point == point {
				foundChild = child
				break
			}
		}

		if foundChild == nil {
			foundChild = NewNode(point, "")
			currentNode.nodes = append(currentNode.nodes, foundChild)
		}
		currentNode = foundChild
	}
	if currentNode.Signatures == nil {
		currentNode.Signatures = make(map[string]*signature.Signature)
	}
	if _, exists := currentNode.Signatures[tSi]; !exists {
		currentNode.Signatures[tSi] = signature.NewSignature(tSi)
	}
}

func (tr *Tree) FindPattern(tSi string) (bool, *Node, int) {
	raw := signature.Unpack(tSi)
	if len(raw) == 0 {
		return false, nil, -1
	}

	var currentNode *Node
	for _, node := range tr.nodes {
		if node.Point == raw[0] {
			currentNode = node
			break
		}
	}
	if currentNode == nil {
		return false, nil, -1
	}

	for i := 1; i < len(raw); i++ {
		found := false
		for _, child := range currentNode.nodes {
			if child.Point == raw[i] {
				currentNode = child
				found = true
				break
			}
		}
		if !found {
			return false, nil, i
		}
	}

	if currentNode.ContainsSi(tSi) {
		return true, currentNode, len(raw)
	}
	return false, currentNode, len(raw)
}

func (tr *Tree) FindFromPosition(source []byte, position int64) (string, error) {
	idx := int(position)
	if idx < 0 || idx >= len(source) {
		return "", errors.New("Position is out of bounds")
	}

	for _, root := range tr.nodes {
		foundSig := tr.recursiveSearch(root, source, idx)
		if foundSig != "" {
			return foundSig, nil
		}
	}

	return "", nil
}

func (tr *Tree) recursiveSearch(node *Node, source []byte, idx int) string {
	currentByteStr := hexTable[source[idx]]
	if node.Point != "??" && node.Point != currentByteStr {
		return ""
	}
	var deepResult string
	if idx+1 < len(source) {
		for _, child := range node.nodes {
			res := tr.recursiveSearch(child, source, idx+1)
			if res != "" {
				deepResult = res
				break
			}
		}
	}
	if deepResult != "" {
		return deepResult
	}
	if len(node.Signatures) > 0 {
		for sigStr := range node.Signatures {
			return sigStr
		}
	}

	return ""
}
func (tr *Tree) LoadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		line = strings.ToLower(line)
		tr.Paste(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
