package lib

import (
	"strconv"
	"strings"
)

type INode interface {
	RegChildren([]INode)
	GetValue() int
	GetBody(...int) string
}

type Node struct {
	children []INode
	Value    int
}

func NewNode(value int) *Node {
	return &Node{
		children: []INode{},
		Value:    value,
	}
}

func (n *Node) RegChildren(children []INode) {
	n.children = children
}

func (n *Node) GetValue() int {
	return n.Value
}

func (n *Node) GetBody(stage ...int) string {
	var body string
	if len(stage) > 0 {
		body = strings.Repeat("---", stage[0])
	} else {
		stage = append(stage, 0)
	}
	body += strconv.Itoa(n.GetValue())
	if len(n.children) > 0 {
		body += ":"
		for _, child := range n.children {
			body += "\n" + child.GetBody(stage[0]+1)
		}
	}
	return body
}
